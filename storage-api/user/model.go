package user

type UserGrade struct {
	UserId        string `json:"user_id"`
	PostpaidLimit int    `json:"postpaid_limit"`
	Spp           *int   `json:"spp,omitempty"`
	ShippingFee   *int   `json:"shipping_fee,omitempty"`
	ReturnFee     *int   `json:"return_fee,omitempty"`
}
