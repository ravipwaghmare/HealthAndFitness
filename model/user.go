package model

// User model in DB
type User struct {
	ID            uint   `json:"user_id" gorm:"primary_key"`
	Title         string `json:"title"`
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	DOB           string `json:"dob"`
	Password      string `json:"password"`
	AccountStatus bool   `json:"accountStatus"`
}
