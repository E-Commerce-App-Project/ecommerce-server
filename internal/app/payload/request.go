package payload

type LoginPayload struct {
	Email    string `json:"email" example:"foo@bar.com" form:"email"`
	Password string `json:"password" example:"secret" form:"password"`
} // @name LoginPayload

type RegisterPayload struct {
	Email    string `json:"email" example:"foo@bar.com" form:"email"`
	Name     string `json:"name" example:"John Doe" form:"name"`
	Phone    string `json:"phone" example:"081234567890" form:"phone"`
	Password string `json:"password" example:"secret" form:"password"`
	Address  string `json:"address" example:"Jl. Jenderal Sudirman No. 1" form:"address"`
} //@name RegisterPayload

type AddProductToCartPayload struct {
	ProductID uint `json:"product_id" example:"1" form:"product_id"`
	Quantity  int  `json:"quantity" example:"1" form:"quantity"`
	UserID    uint `json:"user_id" example:"1" form:"user_id"`
} //@name AddProductToCartPayload

type GetUserCartPayload struct {
	UserID int `json:"user_id" example:"1" form:"user_id"`
} //@name GetUserCartPayload
