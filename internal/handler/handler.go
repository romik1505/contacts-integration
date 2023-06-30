package handler

import (
	"github.com/gorilla/schema"
	contact_service "week3_docker/internal/service/contact"
)

type Handler struct {
	ContactService contact_service.IService
}

var decoder = schema.NewDecoder()

func NewHandler(cs contact_service.IService) *Handler {
	return &Handler{
		ContactService: cs,
	}
}
