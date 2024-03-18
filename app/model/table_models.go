package model

type CreateTableRequest struct {
    Name string `json:"name" form:"name"`
}

type GeneralTableResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
}

type DetailedTableResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Fields DetailedFieldResponse
}
