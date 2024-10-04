package v1

type CreateUserAddressRequest struct {
	RecipientName    string `json:"recipientName" binding:"required"`
	RecipientPhone   string `json:"recipientPhone" binding:"required"`
	RecipientAddress string `json:"recipientAddress" binding:"required"`
	Default          bool   `json:"default"  binding:"required"`
}

type UpdateUserAddressRequest struct {
	ID               uint64 `json:"addressId" binding:"required"`
	RecipientName    string `json:"recipientName" binding:"required"`
	RecipientPhone   string `json:"recipientPhone" binding:"required"`
	RecipientAddress string `json:"recipientAddress" binding:"required"`
	Default          bool   `json:"default"  binding:"required"`
}

type GetUserAddressResponseData struct {
	RecipientName    string `json:"recipientName"`
	RecipientPhone   string `json:"recipientPhone"`
	RecipientAddress string `json:"recipientAddress"`
	Default          bool   `json:"default"`
}
type GetUserAddressResponse struct {
	Response
	Data GetUserAddressResponseData
}
type GetUserAddressesResponse struct {
	Response
	Data []GetUserAddressResponseData
}
