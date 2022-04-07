package payload

type LoginPayload struct {
	Email    string `json:"email" example:"foo@bar.com" form:"email"`
	Password string `json:"password" example:"secret" form:"password"`
} // @name LoginPayload

type RegisterPayload struct {
	Email    string `json:"email" example:"foo@bar.com" form:"email"`
	Name     string `json:"name" example:"John Doe" form:"name"`
	Phone    string `json:"phone_number" example:"081234567890" form:"phone"`
	Password string `json:"password" example:"secret" form:"password"`
	Address  string `json:"address" example:"Jl. Jenderal Sudirman No. 1" form:"address"`
} //@name RegisterPayload

type CreateProductPayload struct {
	Name        string `json:"name" example:"sepatu"`
	Price       int    `json:"price" example:"500000"`
	Description string `json:"description" example:"Sepatu lari"`
	Image       string `json:"image" example:"image"`
	Stock       int    `json:"stock" example:"10"`
	UserID      uint   `json:"user_id" example:"1"`
}
