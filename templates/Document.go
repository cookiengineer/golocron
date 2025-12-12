package templates

import _ "embed"
import "text/template"

var DocumentTemplate *template.Template

//go:embed Document.html
var DocumentBytes []byte

func init() {

	root := template.New("root").Funcs(template.FuncMap{
		"RenderStrings": RenderStrings,
		"RenderDomain":  RenderDomain,
		"RenderURL":     RenderURL,
	})

	template.Must(root.New("_document_").Parse(string(DocumentBytes)))

	DocumentTemplate = root.Lookup("_document_")

}
