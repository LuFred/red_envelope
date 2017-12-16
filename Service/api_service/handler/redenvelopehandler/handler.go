package redenvelopehandler

import (
	redemodel "github.com/lufred/red_envelope/Service/api_service/model/redenvelope"
)

// Create 创建红包
// @Title Create
// @Description  创建红包
// @Accept  json
// @Param   amount    		body  int32    true       "红包金额"
// @Param   count    		body  int32    true       "红包个数"
// @Param   token    		header  string  true      "Authorization token"
// @Success 200 {object} redemodel.CreateResponse  "响应"
// @Failure 401 {object} string "token错误"
// @Failure 400 {object} string "参数错误 101:余额不足"
// @Failure 500 {object} string "服务器内部错误"
// @Router /rede [post]
func (h *RedEnvelopeHandler) Create(dto *redemodel.CreateRequest) {

}

// Take 抢红包
// @Title Take
// @Description  抢红包
// @Accept  json
// @Param   rede_id    		body  int32    true       "红包id"
// @Param   secret_code    	body  string    true       "红包口令"
// @Param   token    		header  string  true      "Authorization token"
// @Success 200 {object} redemodel.TakeResponse  "响应"
// @Failure 401 {object} string "token错误"
// @Failure 400 {object} string "参数错误 100:红包不存在或已过期|101:红包已被抢光"
// @Failure 500 {object} string "服务器内部错误"
// @Router /rede/take [post]
func (h *RedEnvelopeHandler) Take(dto *redemodel.TakeRequest) {
}

// GetBalance 查询余额
// @Title GetBalance
// @Description  查询余额
// @Accept  json
// @Param   token    		header  string  true      "Authorization token"
// @Success 200 {object} redemodel.GetBalanceResponse  "响应"
// @Failure 401 {object} string "token错误"
// @Failure 500 {object} string "服务器内部错误"
// @Router /rede/balance [get]
func (h *RedEnvelopeHandler) GetBalance(dto *redemodel.GetBalanceRequest) {
}

// GetRecord 查询红包记录
// @Title GetRecord
// @Description  查询红包记录
// @Accept  json
// @Param   token    		header  string  true      "Authorization token"
// @Success 200 {object} redemodel.GetRecordResponse  "响应"
// @Failure 401 {object} string "token错误"
// @Failure 500 {object} string "服务器内部错误"
// @Router /rede/record [get]
func (h *RedEnvelopeHandler) GetRecord(dto *redemodel.GetRecordRequest) {
}
