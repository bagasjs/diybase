package model

type CreateUserModel struct {
    ID uint `json:"id"`
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
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
