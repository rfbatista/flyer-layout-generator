// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package generate

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"algvisual/internal/entities"
	"algvisual/web/components/progress"
	"algvisual/web/components/topbar"
	"algvisual/web/views"
	"algvisual/web/views/generate/components/editor"
	"algvisual/web/views/generate/components/editor_form"
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

func Page(props PageProps, css string, js string) templ.Component {
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
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dist/%s?time=%s", css, time.Now().String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/page.templ`, Line: 28, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" rel=\"stylesheet\"></head><body><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.1/cdn/components/dropdown/dropdown.js\"></script><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.1/cdn/components/menu/menu.js\"></script><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.1/cdn/components/menu-item/menu-item.js\"></script><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.1/cdn/components/menu-label/menu-label.js\"></script><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.15.1/cdn/components/divider/divider.js\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = topbar.EditorTopBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"with-sidebar\"><div><div class=\"box\">")
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"box\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = editor.Editor(editor.EditorProps{
			Layout:     props.layout,
			Layoutjson: props.layoutjson,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"box\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = editorform.EditorForm(editorform.EditorFormProps{
			DesignID: props.designID,
			Layout:   props.layout,
			Types:    props.types,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = progress.ProgressSpinner("request-progress").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"/web/js/jquery.multi-select.js\" type=\"text/javascript\"></script><script src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dist/%s?time=%s", js, time.Now().String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/generate/page.templ`, Line: 65, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"text/javascript\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
