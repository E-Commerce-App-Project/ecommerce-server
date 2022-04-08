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

type AddProductToCartPayload struct {
	ProductID int `json:"product_id" example:"1" form:"product_id"`
	Quantity  int `json:"quantity" example:"1" form:"quantity"`
	UserID    int `json:"user_id" example:"1" form:"user_id"`
} //@name AddProductToCartPayload

type UpdateCartPayload struct {
	CartID   int `json:"product_id" example:"1" form:"product_id"`
	Quantity int `json:"quantity" example:"1" form:"quantity"`
	UserID   int `json:"user_id" example:"1" form:"user_id"`
} //@name AddProductToCartPayload

type GetUserCartPayload struct {
	UserID int `json:"user_id" example:"1" form:"user_id"`
} //@name GetUserCartPayload

type DeleteUserCartPayload struct {
	UserID     int `json:"user_id" example:"1" form:"user_id"`
	CartItemID int `json:"cart_item_id" example:"1" form:"cart_item_id"`
} //@name DeleteUserCartPayload

type CheckoutPayload struct {
	UserID      int   `json:"user_id" example:"1" form:"user_id"`
	CartItemIDs []int `json:"cart_item_ids" example:"1,2,3" form:"cart_item_ids"`
}

type TransactionPayload struct {
	UserID     int                      `json:"user_id" example:"1" form:"user_id"`
	TotalPrice int                      `json:"total" example:"1" form:"total"`
	CartItems  []TransactionItemPayload `json:"detail" example:"[{\"product_id\":1,\"quantity\":1,\"price\":1}]" form:"detail"`
}

type TransactionItemPayload struct {
	ProductID int `json:"product_id" example:"1" form:"product_id"`
	Quantity  int `json:"quantity" example:"1" form:"quantity"`
	Price     int `json:"price" example:"1" form:"price"`
}
type CreateProductPayload struct {
	Name        string `json:"name" example:"sepatu"`
	Price       int    `json:"price" example:"500000"`
	Description string `json:"description" example:"Sepatu lari"`
	Image       string `json:"image" example:"image"`
	Stock       int    `json:"stock" example:"10"`
	UserID      uint   `json:"user_id" example:"1"`
}
