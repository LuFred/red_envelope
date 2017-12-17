package redenvelopehandler

import (
	basehandle "github.com/lufred/red_envelope/Service/api_service/handler"
	logiccore "github.com/lufred/red_envelope/Service/api_service/logic_service/core"
	logicservice "github.com/lufred/red_envelope/Service/api_service/logic_service/redenvelopeservice"
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
	if h.Ctx.GetUser() == nil {
		h.UnauthorizedAbort("")
		return
	}
	resp, err := logicservice.Create(h.Ctx.Request.Context(),
		h.Ctx.GetUser(), dto)
	if err != nil {
		logicErr, ok := err.(*logiccore.LogicError)
		if !ok {
			h.ErrorAbort(err.Error())
			return
		}
		switch logicErr.Type {
		default:
			h.ErrorAbort(logicErr.Error())
		case 1:
			errData := &basehandle.BadRequestError{
				Code:    logicErr.Code,
				Message: logicErr.Message,
			}
			h.CustomJsonAbort(400, errData)
		}
		return
	}
	h.Data["json"] = resp
	h.ServeJSON()

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
// @Failure 400 {object} string "参数错误 100:红包不存在或已过期|101:不可重复领取同一个红包|102:红包已被抢光|103:口令错误|200:系统繁忙"
// @Failure 500 {object} string "服务器内部错误"
// @Router /rede/take [post]
func (h *RedEnvelopeHandler) Take(dto *redemodel.TakeRequest) {
	if h.Ctx.GetUser() == nil {
		h.UnauthorizedAbort("")
		return
	}
	resp, err := logicservice.Take(h.Ctx.Request.Context(),
		h.Ctx.GetUser(), dto)
	if err != nil {
		logicErr, ok := err.(*logiccore.LogicError)
		if !ok {
			h.ErrorAbort(err.Error())
			return
		}
		switch logicErr.Type {
		default:
			h.ErrorAbort(logicErr.Error())
		case 1:
			errData := &basehandle.BadRequestError{
				Code:    logicErr.Code,
				Message: logicErr.Message,
			}
			h.CustomJsonAbort(400, errData)
		}
		return
	}
	h.Data["json"] = resp
	h.ServeJSON()
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
	if h.Ctx.GetUser() == nil {
		h.UnauthorizedAbort("")
		return
	}
	resp, err := logicservice.GetBalance(h.Ctx.Request.Context(),
		h.Ctx.GetUser(), dto)
	if err != nil {
		logicErr, ok := err.(*logiccore.LogicError)
		if !ok {
			h.ErrorAbort(err.Error())
			return
		}
		switch logicErr.Type {
		default:
			h.ErrorAbort(logicErr.Error())
		case 1:
			errData := &basehandle.BadRequestError{
				Code:    logicErr.Code,
				Message: logicErr.Message,
			}
			h.CustomJsonAbort(400, errData)
		}
		return
	}
	h.Data["json"] = resp
	h.ServeJSON()
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
	if h.Ctx.GetUser() == nil {
		h.UnauthorizedAbort("")
		return
	}
	resp, err := logicservice.GetRecord(h.Ctx.Request.Context(),
		h.Ctx.GetUser(), dto)
	if err != nil {
		logicErr, ok := err.(*logiccore.LogicError)
		if !ok {
			h.ErrorAbort(err.Error())
			return
		}
		switch logicErr.Type {
		default:
			h.ErrorAbort(logicErr.Error())
		case 1:
			errData := &basehandle.BadRequestError{
				Code:    logicErr.Code,
				Message: logicErr.Message,
			}
			h.CustomJsonAbort(400, errData)
		}
		return
	}
	h.Data["json"] = resp
	h.ServeJSON()
}
