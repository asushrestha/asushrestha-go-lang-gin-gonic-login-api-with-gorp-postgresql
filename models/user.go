package models

type User struct {
	Id        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"user_password" json:"password"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}
