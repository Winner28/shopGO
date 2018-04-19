package resources

import "html/template"

type Templates struct {
	templates map[string]*template.Template
}

var templates *Templates

func init() {
	templates = new(Templates)
	templates.templates = make(map[string]*template.Template)
}

func getTemplatesContainer() *Templates {
	return templates
}

func (tmpl *Templates) GetTemplate(name string) *template.Template {
	return tmpl.templates[name]
}

func (tmpl *Templates) AddTemplate(name string, template *template.Template) {
	tmpl.templates[name] = template
}
