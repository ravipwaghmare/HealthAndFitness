package model

// User model in DB
type User struct {
	ID            uint    `json:"user_id" gorm:"primary_key"`
	UserName      string  `json:"username"`
	Name          string  `json:"name"`
	Gender        string  `json:"gender"`
	Email         string  `json:"email"`
	DOB           string  `json:"dob"`
	Password      string  `json:"password"`
	Weight        int     `json:"Weight"`
	Height        int     `json:"Height"`
	AccountStatus uint    `json:"account_status"`
	BMI           float32 `json:"bmi"`
	WeightGoals   int     `json:"weight_goal"`
	StepsGoals    int     `json:"steps_goal"`
}

// SignInResponse response body
type SignInResponse struct {
	User User `json:"user"`
}
