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

func AiUserGetAll(ctx context.Context, req *api.AiUserGetAllReq) (*api.AiUserGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiUser(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiUserGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiUserGetResp{
			ID: v.ID, UserUID: v.UserUID, UserName: v.UserName, CoverImg: v.CoverImg, OnlineStatus: v.OnlineStatus, AuditStatus: v.AuditStatus, UpdateInfo: v.UpdateInfo, UpdateInfoStatus: v.UpdateInfoStatus, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.AiUserGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiUserGet(ctx context.Context, req *api.AiUserGetReq) (*api.AiUserGetResp, error) {
	record, err := dao.GetAiUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserGetResp{
		ID: record.ID, UserUID: record.UserUID, UserName: record.UserName, CoverImg: record.CoverImg, OnlineStatus: record.OnlineStatus, AuditStatus: record.AuditStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func AiUserUpdate(ctx context.Context, req *api.AiUserUpdateReq) (*api.AiUserUpdateResp, error) {
	updated := &model.AiUser{
		ID: req.ID, UserUID: req.UserUID, UserName: req.UserName, CoverImg: req.CoverImg, OnlineStatus: req.OnlineStatus, AuditStatus: req.AuditStatus, UpdateInfo: req.UpdateInfo, UpdateInfoStatus: req.UpdateInfoStatus, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateAiUser(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserUpdateResp{
		ID: record.ID, UserUID: record.UserUID, UserName: record.UserName, CoverImg: record.CoverImg, OnlineStatus: record.OnlineStatus, AuditStatus: record.AuditStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func AiUserDelete(ctx context.Context, req *api.AiUserDeleteReq) (*api.AiUserDeleteResp, error) {
	_, err := dao.DeleteAiUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserDeleteResp{}
	return result, nil
}

func AiUserOnline(ctx context.Context, req *api.AiUserOnlineReq) (*api.AiUserOnlineResp, error) {
	_, _, err := dao.OnlineAiUser(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserOnlineResp{}
	return result, nil
}
