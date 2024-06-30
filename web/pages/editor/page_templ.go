// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package editor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"algvisual/internal/entities"
	"algvisual/web/views"
	"algvisual/web/views/generate/components/generate_form"
	"fmt"
	"time"
)

type PageProps struct {
	files      []entities.DesignFile
	designID   int32
	template   []entities.Template
	types      []string
	layout     entities.Layout
	layoutjson string
}

func Page(props PageProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = views.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dist/vite/assets/%s?time=%s", "style.css", time.Now().String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/pages/editor/page.templ`, Line: 24, Col: 93}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" rel=\"stylesheet\"><link href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dist/vite/assets/%s?time=%s", "editor.css", time.Now().String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/pages/editor/page.templ`, Line: 25, Col: 94}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" rel=\"stylesheet\"></head><body><main class=\"with-sidebar\"><div><div class=\"box\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = generateform.GenerateForm(generateform.GenerateFormProps{
			DesignID: props.designID,
			Layout:   props.layout,
			Types:    props.types,
			Template: props.template,
			Files:    props.files,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"\"><div id=\"root\"></div></div></div></main><script crossorigin src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dist/vite/assets/%s?time=%s", "editor.js", time.Now().String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/pages/editor/page.templ`, Line: 44, Col: 106}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"module\"></script><script src=\"/web/js/jquery.multi-select.js\" type=\"text/javascript\"></script><script>\n\n$(document).ready(function () {\n  $(\"#priority-list\").sortable();\n  $(\"#gen-btn\").on(\"click\", function (e) {\n    orderPriorities();\n    const form = document.getElementById(\"generate-request-form\");\n    const formData = new FormData(form);\n    $.ajax({\n      url: \"/editor/create/image\",\n      type: \"POST\",\n      data: formData,\n      processData: false, // Prevent jQuery from automatically transforming the data into a query string\n      contentType: false, // Set contentType to false for FormData\n      success: function (result) {\n        console.log(\"foi!\", result);\n        $(\"#canvas-container\").html(result); // this will replace when you select the checkbox\n      },\n    });\n  });\n\n  $(\"#batch-btn\").on(\"click\", function () {\n    orderPriorities();\n    submitForm(\"/request/batch\");\n  });\n\n  function submitForm(action) {\n    const form = document.getElementById(\"generate-request-form\");\n    const formData = new FormData(form);\n\n    fetch(action, {\n      method: \"POST\",\n      body: formData,\n    })\n      .then((response) => {\n        console.log(response);\n      })\n      .then((data) => {\n        console.log(\"Success:\", data);\n      })\n      .catch((error) => {\n        console.error(\"Error:\", error);\n      });\n  }\n});\n\nfunction sort() {\n  return {\n    config: {\n      animation: 150,\n      ghostClass: \"opacity-20\",\n      dragClass: \"bg-blue-50\",\n    },\n    init() {\n      Sortable.create(this.$refs.items, this.config);\n    },\n  };\n}\nconst Toast = Swal.mixin({\n  toast: true,\n  position: \"center\",\n  iconColor: \"white\",\n  customClass: {\n    popup: \"colored-toast\",\n  },\n  showConfirmButton: false,\n  timer: 1500,\n  timerProgressBar: true,\n});\n\ndocument.body.addEventListener(\"makeToast\", async function (evt) {\n  if (evt.detail.level == \"success\") {\n    await Toast.fire({\n      icon: \"success\",\n      title: \"Sucesso\",\n    });\n  } else {\n    await Toast.fire({\n      icon: \"error\",\n      title: evt.detail.message,\n    });\n  }\n});\n\ndocument\n  .getElementById(\"generate-request-form\")\n  .addEventListener(\"submit\", function (event) {\n    Array.from(document.getElementsByClassName(\"priority-item\")).forEach(\n      (i) => {\n        const text = i.querySelector(\"#priority\").textContent;\n        i.querySelector(\"#hiddenInput\").value = text;\n      },\n    );\n  });\n\nfunction orderPriorities() {\n  Array.from(document.getElementsByClassName(\"priority-item\")).forEach((i) => {\n    const text = i.querySelector(\"#priority\").textContent;\n    i.querySelector(\"#hiddenInput\").value = text;\n  });\n}\n\n</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
