---
to: web/views/<%= h.changeCase.snake(name) %>/page.templ
---
package <%= h.changeCase.snake(name) %>

import (
	"algvisual/web/views"
	"time"
  "fmt"
)

type PageProps struct {
}

templ Page(props PageProps, css string, js string) {
  <html>
  <head>
			@views.Header()
			<link href={ fmt.Sprintf("/dist/%s?time=%s", css, time.Now().String()) } rel="stylesheet"/>
  </head>
  <body>
			<script src={ fmt.Sprintf("/dist/%s?time=%s", js, time.Now().String()) } type="text/javascript"></script>
  </body>
  </html>
}
