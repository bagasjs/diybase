package entity

type Field struct {
    ID uint
    Name string
    Type string
    Length uint

    IsNullable bool
    IsUnique bool
    IsPrimaryKey bool
    IsAutoIncrement bool

    IsForeignKey bool
    Reference string
    Related string
    OnDelete string
}
