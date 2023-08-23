package models

type AddRequest struct {
	AuthParams UserAuthParams
	Date       UserDate
}

type UserAuthParams struct {
	Login    string
	Password string
}

type UserDate struct {
	Name string
	Age  int
}
