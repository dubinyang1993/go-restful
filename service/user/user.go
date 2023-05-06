package user

import "dubinyang1993/go-restful/model"

type AddReq struct {
	Name  string
	Phone string
}

func Add(name, phone string) error {
	var user model.User
	user.Name = name
	user.Phone = phone

	err := user.Create()
	if err != nil {
		return err
	}
	return nil
}
