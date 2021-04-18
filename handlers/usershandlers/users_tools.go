package usershandlers

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/ent/users"
	"hello/restapi/operations/user"
	"regexp"
)

func VerifPseudoDBRegister(ctx context.Context, pseudo string, h *usersCreateHandler) (*ent.Users, error) {
	user, err := h.ent.Users.Query().Where(users.PseudoEQ(pseudo)).First(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}

func VerifPseudoDBPseudo(ctx context.Context, pseudo string, h *usersPseudoHandler) (*ent.Users, error) {
	user, err := h.ent.Users.Query().Where(users.PseudoEQ(pseudo)).First(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}

func VerifEmailDB(ctx context.Context, email string, h *usersCreateHandler) (*ent.Users, error) {
	user, err := h.ent.Users.Query().Where(users.EmailEQ(email)).First(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}

func VerifMailFormat(params user.RegisterParams) *user.RegisterParams {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString
	if !re(*params.Signup.Email) {
		return &params
	}
	return nil
}

func VerifNilFormat(params user.RegisterParams) *user.RegisterParams {
	if *params.Signup.Pseudo == "" || *params.Signup.Password == "" || *params.Signup.Email == "" || *params.Signup.Fname == "" || params.Signup.Lname == "" {
		return nil
	}
	return &params
}

func VerifNumberFormat(params user.RegisterParams) *user.RegisterParams {
	isAlpha := regexp.MustCompile(`^[a-z]+$`).MatchString
	for _, username := range []string{*params.Signup.Pseudo, *params.Signup.Fname, params.Signup.Lname} {
		if !isAlpha(username) {
			return &params
		}
	}
	return nil
}
