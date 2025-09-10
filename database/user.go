package database

var userList []User

type User struct {
	ID        	int    `json:"id"`
	FirstName 	string `json:"firstName"`
	LastName  	string `json:"lastName"`
	Email     	string `json:"email"`
	Password  	string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func AddUser(u *User) *User {
	if(u.ID != 0){
		return u
	}

	u.ID = len(userList) + 1
	userList = append(userList, *u)

	return u
}

func GetUsers() []User {
	return userList
}

func GetUserByID(id int) *User {
	for _, user := range userList {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func FindUser(email, pass string) (*User) {
		for _, user := range userList { 
		if user.Email == email && user.Password == pass {
			return &user
		} 
	}
	return nil 
}


func init() {
	userList = []User{
		{
			ID:         1,
			FirstName:  "John",
			LastName:   "Doe",
			Email:      "john.doe@example.com",
			Password:   "password123",
			IsShopOwner: false,
		},
		{
			ID:         2,
			FirstName:  "Jane",
			LastName:   "Doe",
			Email:      "jane.doe@example.com",
			Password:   "password123",
			IsShopOwner: true,
		},
		{
			ID:         3,
			FirstName:  "Bob",
			LastName:   "Smith",
			Email:      "bob.smith@example.com",
			Password:   "password123",
			IsShopOwner: false,
		},
	}
}
