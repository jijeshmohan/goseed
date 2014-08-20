package model

type User struct {
	Id    int64
	Name  string `sql:"size:255"`
	Email string
}

func GetAllUsers() (users []User) {
	db.Find(&users)
	return
}
