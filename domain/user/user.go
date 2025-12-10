package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserList []User

func (l UserList) FindByUsername(username string) *User {
	for _, user := range l {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func (l UserList) FindByUserid(id string) *User {
	for _, user := range l {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
