---
to: web/pages/<%= h.changeCase.snake(name) %>/page.templ
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
			<link href={ fmt.Sprintf("/dist/vite/assets/%s?time=%s", "<%= h.changeCase.snake(name) %>.css", time.Now().String()) } rel="stylesheet"/>
  </head>
  <body>
			<div id="root"></div>
			<script crossorigin src={ fmt.Sprintf("/dist/vite/assets/%s?time=%s", "<%= h.changeCase.snake(name) %>.js", time.Now().String()) } type="module"></script>
  </body>
  </html>
}
