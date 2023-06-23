// Code generated by templ@(devel) DO NOT EDIT.

package testelementattributes

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func important() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`width:100;`)
	templCSSID := templ.CSSID(`important`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}

func unimportant() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`width:50;`)
	templCSSID := templ.CSSID(`unimportant`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}

func render(p person) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var var_2 = []any{important()}
		err = templ.RenderCSSItems(ctx, templBuffer, var_2...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div style=\"width: 100;\"")
		if err != nil {
			return err
		}
		if p.important {
			_, err = templBuffer.WriteString(" class=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_2).String()))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		var_3 := `Important`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		var var_4 = []any{unimportant}
		err = templ.RenderCSSItems(ctx, templBuffer, var_4...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div style=\"width: 100;\"")
		if err != nil {
			return err
		}
		if !p.important {
			_, err = templBuffer.WriteString(" class=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_4).String()))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		var_5 := `Unimportant`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		var var_6 = []any{important}
		err = templ.RenderCSSItems(ctx, templBuffer, var_6...)
		if err != nil {
			return err
		}
		var var_7 = []any{unimportant}
		err = templ.RenderCSSItems(ctx, templBuffer, var_7...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div style=\"width: 100;\"")
		if err != nil {
			return err
		}
		if p.important {
			_, err = templBuffer.WriteString(" class=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_6).String()))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString(" class=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_7).String()))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		var_8 := `Else`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
