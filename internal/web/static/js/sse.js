if (typeof (EventSource) !== "undefined") {
    const source = new EventSource("/sse");
    source.onmessage = function (event) {
    console.log(event)
    };
} else {
    console.log("Sorry, your browser does not support server-sent events...");
}
