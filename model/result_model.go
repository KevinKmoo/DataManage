package model

import "github.com/kmookay/MyBittDataManage/constant"

type Result struct {
	StatusCode string      `json:"statusCode"`
	Msg        string      `json:"msg"`
	RetData    interface{} `json:"retData"`
}

func ResultCommon(statusCode string, msg string, retData interface{}) Result {
	return Result{
		StatusCode: statusCode,
		Msg:        msg,
		RetData:    retData,
	}
}

func ResultSuccess(msg string, retData interface{}) Result {
	return Result{
		StatusCode: constant.STATUS_CODE_OK,
		Msg:        msg,
		RetData:    retData,
	}
}

func ResultFail(msg string, retData interface{}) Result {
	return Result{
		StatusCode: constant.STATUS_CODE_FAIL,
		Msg:        msg,
		RetData:    retData,
	}
}
