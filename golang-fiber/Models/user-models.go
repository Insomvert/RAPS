package Models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *User) TableName() string {
	return "user"
}
