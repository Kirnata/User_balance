package User_balance

type History struct {
	Id       int `json:"-" db:"id"`
	SenderId int `json:"sender_id" binding:"required"`
	GetterId int `json:"getter_id" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
}
