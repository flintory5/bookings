package models

import "github.com/flintory5/bookings/internal/forms"

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]int
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
	Form *forms.Form
}