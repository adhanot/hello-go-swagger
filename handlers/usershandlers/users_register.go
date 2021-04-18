package usershandlers

import (
	"context"
	"hello/models"
	"hello/pkg/ent"
	"hello/restapi/operations/user"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-openapi/runtime/middleware"
)

type usersCreateHandler struct {
	ent *ent.Client
}

func NewUsersCreateHandler(ent *ent.Client) user.RegisterHandler {
	return &usersCreateHandler{ent: ent}
}

// Handle the get entry request
func (h *usersCreateHandler) Handle(params user.RegisterParams) middleware.Responder {
	ifnil := VerifNilFormat(params)
	if ifnil == nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Can't send nil params"}}
	}
	number := VerifNumberFormat(params)
	if number != nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Pseudo, fname and lname must be alphabetic characters [a-z]"}}
	}
	_, err := VerifPseudoDBRegister(params.HTTPRequest.Context(), *params.Signup.Pseudo, h)
	if err == nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Pseudo already exist"}}
	}
	mailFormat := VerifMailFormat(params)
	if mailFormat != nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Email format invalid"}}
	}
	_, err = VerifEmailDB(params.HTTPRequest.Context(), *params.Signup.Email, h)
	if err == nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Email already exist"}}
	}
	_, err = UsersCreateDB(params.HTTPRequest.Context(), h, params)
	if err != nil {
		return &user.RegisterBadRequest{Payload: &models.Error{Success: false, Message: "Database Error"}}
	}
	return &user.RegisterOK{Payload: &models.Success{Success: true, Message: "User Sucessfully Create"}}
}

func UsersCreateDB(ctx context.Context, h *usersCreateHandler, params user.RegisterParams) (*ent.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*params.Signup.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser, err := h.ent.Users.
		Create().
		SetPseudo(*params.Signup.Pseudo).
		SetPassword(string(hashedPassword)).
		SetEmail(*params.Signup.Email).
		SetFname(*params.Signup.Fname).
		SetLname(params.Signup.Lname).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
