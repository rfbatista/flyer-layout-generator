// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"algvisual/internal/entities"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/layout"
	"algvisual/internal/web/components/table"
	"algvisual/internal/web/views"
)

type PageProps struct {
	templates []entities.Template
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container column is-10\"><div class=\"my-2\"><div class=\"card is-hidden1\"><form id=\"form\" hx-encoding=\"multipart/form-data\" hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(shared.PageUploadDesignFile.String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/views/templates/page.templ`, Line: 25, Col: 103}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"none\"><div class=\"field\"><div class=\"control\"><button class=\"button\">Cadastrar</button></div></div></form></div></div></div><div class=\"container\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var4 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
					if !templ_7745c5c3_IsBuffer {
						templ_7745c5c3_Buffer = templ.GetBuffer()
						defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
					}
					templ_7745c5c3_Var6 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						templ_7745c5c3_Err = table.THeadCol("Nome").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = table.THeadCol("Largura").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = table.THeadCol("Altura").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = table.THeadCol("Limite de regioes em x").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = table.THeadCol("Limite de regioes em y").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = table.THeadCol("Criado em").Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return templ_7745c5c3_Err
					})
					templ_7745c5c3_Err = table.THeadRow().Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if !templ_7745c5c3_IsBuffer {
						_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
					}
					return templ_7745c5c3_Err
				})
				templ_7745c5c3_Err = table.THead().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Var7 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
					if !templ_7745c5c3_IsBuffer {
						templ_7745c5c3_Buffer = templ.GetBuffer()
						defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
					}
					for _, d := range props.templates {
						templ_7745c5c3_Var8 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							templ_7745c5c3_Err = table.TBodyCol(d.Name).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = table.TBodyCol(d.WidthS()).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = table.TBodyCol(d.HeightS()).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = table.TBodyCol(d.MaxSlotsXText()).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = table.TBodyCol(d.MaxSlotsYText()).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = table.TBodyCol(d.CreatedAtText()).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = table.TBodyRow().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					if !templ_7745c5c3_IsBuffer {
						_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
					}
					return templ_7745c5c3_Err
				})
				templ_7745c5c3_Err = table.TBody().Render(templ.WithChildren(ctx, templ_7745c5c3_Var7), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = table.Table().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n        htmx.on('#form', 'htmx:xhr:progress', function(evt) {\n          htmx.find('#progress').setAttribute('value', evt.detail.loaded/evt.detail.total * 100)\n        });\n        const Toast = Swal.mixin({\n          toast: true,\n          position: 'center',\n          iconColor: 'white',\n          customClass: {\n            popup: 'colored-toast',\n          },\n          showConfirmButton: false,\n          timer: 1500,\n          timerProgressBar: true,\n        })\n        document.body.addEventListener(\"makeToast\", async function(evt){\n          console.log(evt.detail)\n          if(evt.detail.level == \"success\"){\n            await Toast.fire({\n              icon: 'success',\n              title: 'Sucesso',\n            })\n          } else {\n            await Toast.fire({\n              icon: 'error',\n              title: evt.detail.message,\n            })\n          }\n        })\n        </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
