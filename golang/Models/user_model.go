package Models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (s *User) TableName() string {
	return "user" // database table name for this model
}
