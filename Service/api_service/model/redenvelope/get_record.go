package redenvelope

//GetRecordRequest 查看余额request
type GetRecordRequest struct{}

//GetRecordResponse 查看余额response
type GetRecordResponse struct {
	Data []*Record `json:"data" description:"红包记录集合"`
}

//Record 红包记录dto
type Record struct {
	RedeID int32 `json:"rede_id" description:"红包id"`
	Amount int32 `json:"amount" description:"抢到红包金额(精确到分的整数)"`
	Time   int64 `json:"time" description:"红包领取时间(毫秒时间戳)"`
}
