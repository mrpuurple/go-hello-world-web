package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
<<<<<<< HEAD
	IntMap map[string]int
=======
	IntMap map[int]int
>>>>>>> 888a37c (refactor)
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}