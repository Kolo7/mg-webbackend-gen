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

func AiCreditActInfoGetAll(ctx context.Context, req *api.AiCreditActInfoGetAllReq) (*api.AiCreditActInfoGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiCreditActInfo(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiCreditActInfoGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiCreditActInfoGetResp{
			ID: v.ID, ActID: v.ActID, UserID: v.UserID, GiftID: v.GiftID, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt,
		})
	}

	result := &api.AiCreditActInfoGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiCreditActInfoGet(ctx context.Context, req *api.AiCreditActInfoGetReq) (*api.AiCreditActInfoGetResp, error) {
	record, err := dao.GetAiCreditActInfo(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActInfoGetResp{
		ID: record.ID, ActID: record.ActID, UserID: record.UserID, GiftID: record.GiftID, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}

	return result, nil
}

func AiCreditActInfoUpdate(ctx context.Context, req *api.AiCreditActInfoUpdateReq) (*api.AiCreditActInfoUpdateResp, error) {
	updated := &model.AiCreditActInfo{
		ID: req.ID, ActID: req.ActID, UserID: req.UserID, GiftID: req.GiftID, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt,
	}
	record, _, err := dao.UpdateAiCreditActInfo(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActInfoUpdateResp{
		ID: record.ID, ActID: record.ActID, UserID: record.UserID, GiftID: record.GiftID, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}
	return result, nil
}

func AiCreditActInfoDelete(ctx context.Context, req *api.AiCreditActInfoDeleteReq) (*api.AiCreditActInfoDeleteResp, error) {
	_, err := dao.DeleteAiCreditActInfo(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditActInfoDeleteResp{}
	return result, nil
}
