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

func AiCreditGiftGetAll(ctx context.Context, req *api.AiCreditGiftGetAllReq) (*api.AiCreditGiftGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiCreditGift(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiCreditGiftGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiCreditGiftGetResp{
			ID: v.ID, GiftID: v.GiftID, GiftName: v.GiftName, GiftImg: v.GiftImg, CreditLevel: v.CreditLevel, GiftURL: v.GiftURL, GiftTips: v.GiftTips, IsEnd: v.IsEnd, OnlineStatus: v.OnlineStatus, AuditStatus: v.AuditStatus, UpdateInfo: v.UpdateInfo, UpdateInfoStatus: v.UpdateInfoStatus, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.AiCreditGiftGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiCreditGiftGet(ctx context.Context, req *api.AiCreditGiftGetReq) (*api.AiCreditGiftGetResp, error) {
	record, err := dao.GetAiCreditGift(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditGiftGetResp{
		ID: record.ID, GiftID: record.GiftID, GiftName: record.GiftName, GiftImg: record.GiftImg, CreditLevel: record.CreditLevel, GiftURL: record.GiftURL, GiftTips: record.GiftTips, IsEnd: record.IsEnd, OnlineStatus: record.OnlineStatus, AuditStatus: record.AuditStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func AiCreditGiftUpdate(ctx context.Context, req *api.AiCreditGiftUpdateReq) (*api.AiCreditGiftUpdateResp, error) {
	updated := &model.AiCreditGift{
		ID: req.ID, GiftID: req.GiftID, GiftName: req.GiftName, GiftImg: req.GiftImg, CreditLevel: req.CreditLevel, GiftURL: req.GiftURL, GiftTips: req.GiftTips, IsEnd: req.IsEnd, OnlineStatus: req.OnlineStatus, AuditStatus: req.AuditStatus, UpdateInfo: req.UpdateInfo, UpdateInfoStatus: req.UpdateInfoStatus, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateAiCreditGift(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditGiftUpdateResp{
		ID: record.ID, GiftID: record.GiftID, GiftName: record.GiftName, GiftImg: record.GiftImg, CreditLevel: record.CreditLevel, GiftURL: record.GiftURL, GiftTips: record.GiftTips, IsEnd: record.IsEnd, OnlineStatus: record.OnlineStatus, AuditStatus: record.AuditStatus, UpdateInfo: record.UpdateInfo, UpdateInfoStatus: record.UpdateInfoStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func AiCreditGiftDelete(ctx context.Context, req *api.AiCreditGiftDeleteReq) (*api.AiCreditGiftDeleteResp, error) {
	_, err := dao.DeleteAiCreditGift(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditGiftDeleteResp{}
	return result, nil
}

func AiCreditGiftOnline(ctx context.Context, req *api.AiCreditGiftOnlineReq) (*api.AiCreditGiftOnlineResp, error) {
	_, _, err := dao.OnlineAiCreditGift(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditGiftOnlineResp{}
	return result, nil
}
