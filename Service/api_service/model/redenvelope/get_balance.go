package redenvelope

//GetBalanceRequest 查看余额request
type GetBalanceRequest struct{}

//GetBalanceResponse 查看余额response
type GetBalanceResponse struct {
	Balance int32 `json:"balance" description:"余额"`
}
