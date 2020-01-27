package Models

type Token struct {
	UserId uint
	jwt.StandardClaims

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}