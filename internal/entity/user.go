package entity

type UserTeste struct {
	ID   int
	Name string
	Age  int
}

func (UserTeste) TableName() string {
	return "userteste"
}