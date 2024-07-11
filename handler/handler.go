package handler

import "text/template"

type Handler struct {
	templateIndex *template.Template
}

func New(
	templateIndex *template.Template,
) *Handler {
	return &Handler{
		templateIndex: templateIndex,
	}
}
