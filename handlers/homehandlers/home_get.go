package homehandlers

import (
	"hello/models"
	"hello/restapi/operations/home"

	"github.com/go-openapi/runtime/middleware"
)

// NewHomeHandler handles a request for getting an entry
func NewHomeHandler() home.GetHelloHandler {
	return &getHome{}
}

type getHome struct {
}

// Handle the get entry request
func (h *getHome) Handle(params home.GetHelloParams, jwt interface{}) middleware.Responder {
	return &home.GetHelloOK{Payload: &models.WelcomingMessage{Message: "Welcome on Automate !"}}
}
