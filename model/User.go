package model

import "time"

type DataUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type DataRegister struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Role     string `json:"role" bson:"role" validate:"required"`
}
type ResponseGlobal struct {
	Status         string      `json:"status"`
	StatusDesc     string      `json:"status_desc"`
	StatusDateTime time.Time   `json:"status_date_time"`
	Result         interface{} `json:"result"`
}
type ResponseLogin struct {
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
	Role  string `json:"role" bson:"role"`
	Token string `json:"token"`
}
type DeleteUser struct {
	Username string `json:"username" validate:"required"`
}
type ResponseCheckDataUser struct {
	Username string `json:"username" bson:"username" `
	Email    string `json:"email" bson:"email" `
	Name     string `json:"name" bson:"name" `
	Role     string `json:"role" bson:"role" `
}
type DataUpdate struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Role           string `json:"role"`
	UsernameTarget string `json:"username_target"`
}
