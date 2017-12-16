package redenvelope

//TakeRequest 领取红包Request
type TakeRequest struct {
	RedeID     int32  `json:"rede_id" required:"true" description:"红包id"`
	SecretCode string `json:"secret_code" required:"true" description:"红包口令"`
}

//TakeResponse 领取红包response
type TakeResponse struct {
	Amount int32 `json:"amount" description:"抢到的金额数"`
}
