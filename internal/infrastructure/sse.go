package infrastructure

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewEvent(message string) Event {
	return Event{Data: []byte(message)}
}

// Event represents Server-Sent Event.
// SSE explanation: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format
type Event struct {
	// ID is used to set the EventSource object's last event ID value.
	ID []byte
	// Data field is for the message. When the EventSource receives multiple consecutive lines
	// that begin with data:, it concatenates them, inserting a newline character between each one.
	// Trailing newlines are removed.
	Data []byte
	// Event is a string identifying the type of event described. If this is specified, an event
	// will be dispatched on the browser to the listener for the specified event name; the website
	// source code should use addEventListener() to listen for named events. The onmessage handler
	// is called if no event name is specified for a message.
	Event []byte
	// Retry is the reconnection time. If the connection to the server is lost, the browser will
	// wait for the specified time before attempting to reconnect. This must be an integer, specifying
	// the reconnection time in milliseconds. If a non-integer value is specified, the field is ignored.
	Retry []byte
	// Comment line can be used to prevent connections from timing out; a server can send a comment
	// periodically to keep the connection alive.
	Comment []byte
}

// MarshalTo marshals Event to given Writer
func (ev *Event) MarshalTo(w io.Writer) error {
	// Marshalling part is taken from: https://github.com/r3labs/sse/blob/c6d5381ee3ca63828b321c16baa008fd6c0b4564/http.go#L16
	if len(ev.Data) == 0 && len(ev.Comment) == 0 {
		return nil
	}

	if len(ev.Data) > 0 {
		if _, err := fmt.Fprintf(w, "id: %s\n", ev.ID); err != nil {
			return err
		}

		sd := bytes.Split(ev.Data, []byte("\n"))
		for i := range sd {
			if _, err := fmt.Fprintf(w, "data: %s\n", sd[i]); err != nil {
				return err
			}
		}

		if len(ev.Event) > 0 {
			if _, err := fmt.Fprintf(w, "event: %s\n", ev.Event); err != nil {
				return err
			}
		}

		if len(ev.Retry) > 0 {
			if _, err := fmt.Fprintf(w, "retry: %s\n", ev.Retry); err != nil {
				return err
			}
		}
	}

	if len(ev.Comment) > 0 {
		if _, err := fmt.Fprintf(w, ": %s\n", ev.Comment); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\n"); err != nil {
		return err
	}

	return nil
}

func NewServerSideEventManager(log *zap.Logger) *ServerSideEventManager {
	c := make(chan Event, 5)
	return &ServerSideEventManager{
		log:            log,
		rec:            c,
		Notifier:       make(chan Event, 1),
		newClients:     make(chan chan Event),
		closingClients: make(chan chan Event),
		clients:        make(map[chan Event]bool),
	}
}

type ServerSideEventManager struct {
	rec chan Event
	log *zap.Logger
	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan Event

	// New client connections are pushed to this channel
	newClients chan chan Event

	// Closed client connections are pushed to this channel
	closingClients chan chan Event

	// Client connections registry
	clients map[chan Event]bool
}

func (se *ServerSideEventManager) Listen() {
	for {
		select {
		case s := <-se.newClients:

			// A new client has connected.
			// Register their message channel
			se.clients[s] = true
			se.log.Info(fmt.Sprintf("Client added. %d registered clients", len(se.clients)))
		case s := <-se.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(se.clients, s)
			se.log.Info(fmt.Sprintf("Removed client. %d registered clients", len(se.clients)))
		case event := <-se.Notifier:
			se.log.Debug("sending notification to sse clients")
			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan := range se.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (s ServerSideEventManager) BroadCastEvent(e Event) error {
	// Send the message to the broker via Notifier channel
	s.Notifier <- e
	return nil
}

func (s ServerSideEventManager) HandleConnection(c echo.Context) error {
	s.log.Info("SSE client connected, ip: %v", zap.String("ip", c.RealIP()))
	w := c.Response()
	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan Event)

	// Signal the broker that we have a new connection
	s.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		s.closingClients <- messageChan
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request().Context().Done():
			s.log.Info(fmt.Sprintf("SSE client disconnected, ip: %v", c.RealIP()))
			s.closingClients <- messageChan
			return nil
		case event := <-messageChan:
			s.log.Debug("sending new message to sse client")
			if err := event.MarshalTo(w); err != nil {
				return err
			}
			w.Flush()
		case <-ticker.C:
			event := Event{
				Data: []byte("time: " + time.Now().Format(time.RFC3339Nano)),
			}
			if err := event.MarshalTo(w); err != nil {
				return err
			}
			w.Flush()
		}
	}
}
