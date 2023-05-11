package models

type User struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Address		string		`json:"address"`
	Email		string		`json:"email" gorm:"unique"`
	Password	string		`json:"password"`
}
