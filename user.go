package dairy

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding: "required"`
	Surname  string `json:"surname"`
	Login    string `json:"login" binding: "required"`
	Password string `json:"password" binding: "required"`
}
