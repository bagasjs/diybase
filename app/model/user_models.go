package model

type CreateUserModel struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type UpdateUserModel struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type GeneralUserModel struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
