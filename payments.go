package User_balance

type History struct {
	Id       int `json:"-" db:"id"`
	SenderId int `json:"sender_id"`
	GetterId int `json:"getter_id" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
}

type SinglePay struct {
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
}