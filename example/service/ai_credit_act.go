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

func AiCreditActGetAll(ctx context.Context, req *api.AiCreditActGetAllReq) (*api.AiCreditActGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiCreditAct(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiCreditActGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiCreditActGetResp{
			ID: v.ID, ActName: v.ActName, StartTime: v.StartTime, EndTime: v.EndTime, AuditStatus: v.AuditStatus, OnlineStatus: v.OnlineStatus, UpdateInfo: v.UpdateInfo, UpdateInfoStatus: v.UpdateInfoStatus, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.AiCreditActGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiCreditActGet(ctx context.Context, req *api.AiCreditActGetReq) (*api.AiCreditActGetResp, error) {
	record, err := dao.GetAiCreditAct(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActGetResp{
		ID: record.ID, ActName: record.ActName, StartTime: record.StartTime, EndTime: record.EndTime, AuditStatus: record.AuditStatus, OnlineStatus: record.OnlineStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func AiCreditActUpdate(ctx context.Context, req *api.AiCreditActUpdateReq) (*api.AiCreditActUpdateResp, error) {
	updated := &model.AiCreditAct{
		ID: req.ID, ActName: req.ActName, StartTime: req.StartTime, EndTime: req.EndTime, AuditStatus: req.AuditStatus, OnlineStatus: req.OnlineStatus, UpdateInfo: req.UpdateInfo, UpdateInfoStatus: req.UpdateInfoStatus, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateAiCreditAct(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActUpdateResp{
		ID: record.ID, ActName: record.ActName, StartTime: record.StartTime, EndTime: record.EndTime, AuditStatus: record.AuditStatus, OnlineStatus: record.OnlineStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func AiCreditActDelete(ctx context.Context, req *api.AiCreditActDeleteReq) (*api.AiCreditActDeleteResp, error) {
	_, err := dao.DeleteAiCreditAct(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActDeleteResp{}
	return result, nil
}

func AiCreditActOnline(ctx context.Context, req *api.AiCreditActOnlineReq) (*api.AiCreditActOnlineResp, error) {
	_, _, err := dao.OnlineAiCreditAct(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActOnlineResp{}
	return result, nil
}
