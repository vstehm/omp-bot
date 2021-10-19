package user

import "fmt"

type User struct {
	Id        uint64 `json:"id,omitempty"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("id: %v\nlastname: %v\nfirstname: %v\nphone: %v\nemail: %v\n", u.Id, u.Lastname, u.Firstname, u.Phone, u.Email)
}
