package usershandlers

import (
	"hello/models"
	"hello/pkg/ent"
	"hello/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

type usersPseudoHandler struct {
	ent *ent.Client
}

func NewUsersPseudoHandler(ent *ent.Client) user.LoginHandler {
	return &usersPseudoHandler{ent: ent}
}

// Handle the get entry request
func (h *usersPseudoHandler) Handle(params user.LoginParams) middleware.Responder {
	data, err := VerifPseudoDBPseudo(params.HTTPRequest.Context(), *params.Login.Pseudo, h)
	if err != nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "pseudo not found"}}
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(*params.Login.Password))
	if err != nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Bad Password"}}
	}
	token, err := GenerateJWT(data.Pseudo, data.Email, data.Fname, data.Lname)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return &user.LoginOK{Payload: &models.LoginSuccess{Success: true, Token: token}}
}
