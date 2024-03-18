package entity

type Field struct {
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
