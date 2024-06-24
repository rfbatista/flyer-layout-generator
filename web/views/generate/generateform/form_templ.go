// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package generateform

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"
import "algvisual/internal/entities"

type GenerateFormProps struct {
	Files      []entities.DesignFile
	DesignID   int32
	Template   []entities.Template
	Types      []string
	Layout     entities.Layout
	Layoutjson string
}

func GenerateForm(props GenerateFormProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"generate-request-form\" class=\"form\" hx-indicator=\"#request-progress\" hx-swap=\"none\"><input name=\"layout_id\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", props.Layout.ID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 22, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"display:none;\"> <input name=\"grid_x\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("15")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 23, Col: 35}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"display:none;\"> <input name=\"grid_y\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("15")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 24, Col: 35}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"display:none;\"><div class=\"stack\"><div class=\"stack generation-form__files\"><input name=\"design_id\" class=\"input-style\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", props.DesignID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 27, Col: 89}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", props.DesignID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 27, Col: 139}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"stack generation-form__formatos\"><select class=\"uk-select\" label=\"Formatos\" class=\"\" multiple=\"multiple\" id=\"templates\" name=\"templates[]\" required>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, t := range props.Template {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option type=\"checkbox\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(t.SID())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 32, Col: 45}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(t.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 33, Col: 15}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <i class=\"text-slate-300\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(t.SWidth())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 33, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("x")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(t.SHeigth())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 33, Col: 72}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i></option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"switcher generation-form__formatos__button\"><div><sl-button size=\"small\" type=\"button\" onclick=\"$(&#39;#templates&#39;).multiSelect(&#39;select_all&#39;)\">Selecionar todos</sl-button> <sl-button size=\"small\" type=\"button\" onclick=\"$(&#39;#templates&#39;).multiSelect(&#39;deselect_all&#39;);\">Deselecionar todos</sl-button></div></div></div></div><sl-divider></sl-divider><div class=\"stack generation-form__priorities\"><label class=\"form-label\">Prioridades</label><div class=\"\" x-data=\"sort()\" x-init=\"init()\"><div class=\"stack\" x-ref=\"items\" id=\"items\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, t := range props.Types {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"priority-item with-icon generation-form__priority-list__item\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"icon lucide lucide-grip-vertical\"><circle cx=\"9\" cy=\"12\" r=\"1\"></circle><circle cx=\"9\" cy=\"5\" r=\"1\"></circle><circle cx=\"9\" cy=\"19\" r=\"1\"></circle><circle cx=\"15\" cy=\"12\" r=\"1\"></circle><circle cx=\"15\" cy=\"5\" r=\"1\"></circle><circle cx=\"15\" cy=\"19\" r=\"1\"></circle></svg> <input type=\"hidden\" id=\"hiddenInput\" name=\"priority[]\"> <span class=\"\" id=\"priority\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(t)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/generateform/form.templ`, Line: 65, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><div class=\"generation-form__submit\"><sl-button id=\"gen-btn\" variant=\"primary\" type=\"button\" class=\"btn\">Gerar</sl-button> <sl-button id=\"batch-btn\" variant=\"default\" type=\"button\" class=\"btn\">Gerar Lote</sl-button></div></form><section id=\"toast-container\"></section><script type=\"text/javascript\">\n    document.getElementById('generate-request-form').addEventListener('submit', function(event) {\n        Array.from(document.getElementsByClassName('priority-item')).forEach((i) => {\n         const text = i.querySelector('#priority').textContent;\n          i.querySelector('#hiddenInput').value = text;\n        })\n    });\n\nfunction orderPriorities(){\n      Array.from(document.getElementsByClassName('priority-item')).forEach((i) => {\n       const text = i.querySelector('#priority').textContent;\n        i.querySelector('#hiddenInput').value = text;\n      })\n}\n\n$(document).ready(function () {\n      $('#gen-btn').on('click', function(e) {\n        orderPriorities()\n          const form = document.getElementById('generate-request-form');\n          const formData = new FormData(form);\n      $.ajax({\n            url: '/editor/create/image',\n            type: 'POST',\n            data: formData,\n                    processData: false, // Prevent jQuery from automatically transforming the data into a query string\n                    contentType: false, // Set contentType to false for FormData\n            success: function(result) {\n              console.log(\"foi!\", result)\n              $(\"#canvas-container\").html(result); // this will replace when you select the checkbox\n            }\n        });\n      });\n\n      $('#batch-btn').on('click', function() {\n        orderPriorities()\n        submitForm('/request/batch');\n      });\n\n      function submitForm(action) {\n          const form = document.getElementById('generate-request-form');\n          const formData = new FormData(form);\n          \n          fetch(action, {\n              method: 'POST',\n              body: formData\n          })\n          .then(response => {\n              console.log(response)\n            })\n          .then(data => {\n              console.log('Success:', data);\n          })\n          .catch((error) => {\n              console.error('Error:', error);\n          });\n      }\n})\n\n    function sort() {\n        return {\n            config: {\n                animation: 150,\n                ghostClass: 'opacity-20',\n                dragClass: 'bg-blue-50',\n            },\n            init() {\n                Sortable.create(this.$refs.items, this.config);\n            }\n        }\n    }\n    const Toast = Swal.mixin({\n      toast: true,\n      position: 'center',\n      iconColor: 'white',\n      customClass: {\n        popup: 'colored-toast',\n      },\n      showConfirmButton: false,\n      timer: 1500,\n      timerProgressBar: true,\n    })\n    document.body.addEventListener(\"makeToast\", async function(evt){\n      if(evt.detail.level == \"success\"){\n        await Toast.fire({\n          icon: 'success',\n          title: 'Sucesso',\n        })\n      } else {\n        await Toast.fire({\n          icon: 'error',\n          title: evt.detail.message,\n        })\n      }\n    })\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
