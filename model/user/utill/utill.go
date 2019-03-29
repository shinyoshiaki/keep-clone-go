package utill

import (
	"keep-server/model/user"
)

var (
	db = user.Connect()
)

// UserNum .
func UserNum() int {
	count := 0
	var u []user.User
	db.Find(&u).Count(&count)
	return count
}

func Code2Name(code string) string {
	var u user.User
	db.Find(&u, "Code = ?", code)
	return u.Name
}
