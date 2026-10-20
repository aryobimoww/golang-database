package entity

type Comment struct {
	Id      int32  `db:"id"`
	Email   string `db:"email"`
	Comment string `db:"comment"`
}