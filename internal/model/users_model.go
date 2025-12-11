package model

// user struct holds user information
type User struct {
	ID    int64  `json:"id"`    // user id
	Name  string `json:"name"`  // user name
	Email string `json:"email"` // user email
}
