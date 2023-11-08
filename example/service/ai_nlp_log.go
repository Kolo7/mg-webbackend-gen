package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/kolo7/mg-webbackend-gen/example/api"
	"github.com/kolo7/mg-webbackend-gen/example/dao"
	"github.com/kolo7/mg-webbackend-gen/example/model"

	"github.com/google/uuid"
)

var (
	_ = time.Second

	_ = uuid.UUID{}
	_ sql.Result
)

func AiNlpLogGetAll(ctx context.Context, req *api.AiNlpLogGetAllReq) (*api.AiNlpLogGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiNlpLog(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiNlpLogGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiNlpLogGetResp{
			ID: v.ID, MsgID: v.MsgID, ReqContent: v.ReqContent, AiRet: v.AiRet, AiRetErr: v.AiRetErr, CallbackCode: v.CallbackCode, CallbackRet: v.CallbackRet, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt, SendID: v.SendID, AiUID: v.AiUID,
		})
	}

	result := &api.AiNlpLogGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiNlpLogGet(ctx context.Context, req *api.AiNlpLogGetReq) (*api.AiNlpLogGetResp, error) {
	record, err := dao.GetAiNlpLog(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiNlpLogGetResp{
		ID: record.ID, MsgID: record.MsgID, ReqContent: record.ReqContent, AiRet: record.AiRet, AiRetErr: record.AiRetErr, CallbackCode: record.CallbackCode, CallbackRet: record.CallbackRet, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt, SendID: record.SendID, AiUID: record.AiUID,
	}

	return result, nil
}

func AiNlpLogUpdate(ctx context.Context, req *api.AiNlpLogUpdateReq) (*api.AiNlpLogUpdateResp, error) {
	updated := &model.AiNlpLog{
		ID: req.ID, MsgID: req.MsgID, ReqContent: req.ReqContent, AiRet: req.AiRet, AiRetErr: req.AiRetErr, CallbackCode: req.CallbackCode, CallbackRet: req.CallbackRet, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt, SendID: req.SendID, AiUID: req.AiUID,
	}
	record, _, err := dao.UpdateAiNlpLog(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiNlpLogUpdateResp{
		ID: record.ID, MsgID: record.MsgID, ReqContent: record.ReqContent, AiRet: record.AiRet, AiRetErr: record.AiRetErr, CallbackCode: record.CallbackCode, CallbackRet: record.CallbackRet, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt, SendID: record.SendID, AiUID: record.AiUID,
	}
	return result, nil
}

func AiNlpLogDelete(ctx context.Context, req *api.AiNlpLogDeleteReq) (*api.AiNlpLogDeleteResp, error) {
	_, err := dao.DeleteAiNlpLog(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiNlpLogDeleteResp{}
	return result, nil
}
