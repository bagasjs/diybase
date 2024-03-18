package model

type CreateUserRequest struct {
    Name string `json:"name" form:"name"`
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
    PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
}

type UpdateUserRequest struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type GeneralUserResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
