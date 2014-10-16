package model

type User struct {
	Id    int64
	Name  string `sql:"not null;unique;size:255"`
	Email string
}

func GetAllUsers() (users []User) {
	users = make([]User, 0)
	db.Find(&users)
	return
}

func CreateNewUser(user *User) error {
	result := db.Create(user)
	return result.Error
}
