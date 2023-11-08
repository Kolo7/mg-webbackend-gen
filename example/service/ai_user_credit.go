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

func AiUserCreditGetAll(ctx context.Context, req *api.AiUserCreditGetAllReq) (*api.AiUserCreditGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiUserCredit(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiUserCreditGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiUserCreditGetResp{
			ID: v.ID, UserUID: v.UserUID, AiUID: v.AiUID, ActID: v.ActID, UserCredit: v.UserCredit, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.AiUserCreditGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiUserCreditGet(ctx context.Context, req *api.AiUserCreditGetReq) (*api.AiUserCreditGetResp, error) {
	record, err := dao.GetAiUserCredit(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserCreditGetResp{
		ID: record.ID, UserUID: record.UserUID, AiUID: record.AiUID, ActID: record.ActID, UserCredit: record.UserCredit, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func AiUserCreditUpdate(ctx context.Context, req *api.AiUserCreditUpdateReq) (*api.AiUserCreditUpdateResp, error) {
	updated := &model.AiUserCredit{
		ID: req.ID, UserUID: req.UserUID, AiUID: req.AiUID, ActID: req.ActID, UserCredit: req.UserCredit, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateAiUserCredit(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserCreditUpdateResp{
		ID: record.ID, UserUID: record.UserUID, AiUID: record.AiUID, ActID: record.ActID, UserCredit: record.UserCredit, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func AiUserCreditDelete(ctx context.Context, req *api.AiUserCreditDeleteReq) (*api.AiUserCreditDeleteResp, error) {
	_, err := dao.DeleteAiUserCredit(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserCreditDeleteResp{}
	return result, nil
}
