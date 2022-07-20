package user

var (
	CreateUser        = "INSERT INTO accounts (username, email, password) VALUES ($1,$2,$3);"
	GetUserbyEmail    = "SELECT * from accounts WHERE email = $1 ;"
	GetUserbyusername = "SELECT * from accounts WHERE username = $1 ;"
)
