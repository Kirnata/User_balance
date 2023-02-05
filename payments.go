package User_balance

type History struct {
	Id          int     `json:"-" db:"id"`
	SenderId    int     `json:"sender_id" db:"sender_id"`
	GetterId    int     `json:"getter_id" binding:"required" db:"getter_id"`
	Amount      float32 `json:"amount" binding:"required" db:"value"`
	Description string  `json:"description" binding:"required" db:"description"`
}

type SinglePay struct {
	Amount      float32 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type B2BPay struct {
	GetterId    int     `json:"getter_id" binding:"required"`
	Amount      float32 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
