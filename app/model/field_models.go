package model

type CreateFieldRequest struct {
    ID uint
    Name string
    Type string
    Length uint

    Nullable bool
    Unique bool
    PrimaryKey bool
    AutoIncrement bool

    ForeignKey bool
    Reference string
    Related string
    OnDelete string
}
type GeneralFieldResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
}

type DetailedFieldResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Length uint `json:"length"`

    Nullable bool `json:"nullable"`
    Unique bool `json:"unique"`
    PrimaryKey bool `json:"primary_key"`
    AutoIncrement bool `json:"auto_increment"`

    ForeignKey bool `json:"foreign_key"`
    Reference string `json:"reference"`
    Related string `json:"related"`
    OnDelete string `json:"on_delete"`
}
