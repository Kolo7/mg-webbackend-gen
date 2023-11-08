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

func AiCreditLevelConfGetAll(ctx context.Context, req *api.AiCreditLevelConfGetAllReq) (*api.AiCreditLevelConfGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiCreditLevelConf(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiCreditLevelConfGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiCreditLevelConfGetResp{
			ID: v.ID, CreditLevel: v.CreditLevel, LevelName: v.LevelName, LevelScore: v.LevelScore, LogoURL: v.LogoURL, OnlineStatus: v.OnlineStatus, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt,
		})
	}

	result := &api.AiCreditLevelConfGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiCreditLevelConfGet(ctx context.Context, req *api.AiCreditLevelConfGetReq) (*api.AiCreditLevelConfGetResp, error) {
	record, err := dao.GetAiCreditLevelConf(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditLevelConfGetResp{
		ID: record.ID, CreditLevel: record.CreditLevel, LevelName: record.LevelName, LevelScore: record.LevelScore, LogoURL: record.LogoURL, OnlineStatus: record.OnlineStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}

	return result, nil
}

func AiCreditLevelConfUpdate(ctx context.Context, req *api.AiCreditLevelConfUpdateReq) (*api.AiCreditLevelConfUpdateResp, error) {
	updated := &model.AiCreditLevelConf{
		ID: req.ID, CreditLevel: req.CreditLevel, LevelName: req.LevelName, LevelScore: req.LevelScore, LogoURL: req.LogoURL, OnlineStatus: req.OnlineStatus, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt,
	}
	record, _, err := dao.UpdateAiCreditLevelConf(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditLevelConfUpdateResp{
		ID: record.ID, CreditLevel: record.CreditLevel, LevelName: record.LevelName, LevelScore: record.LevelScore, LogoURL: record.LogoURL, OnlineStatus: record.OnlineStatus, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}
	return result, nil
}

func AiCreditLevelConfDelete(ctx context.Context, req *api.AiCreditLevelConfDeleteReq) (*api.AiCreditLevelConfDeleteResp, error) {
	_, err := dao.DeleteAiCreditLevelConf(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditLevelConfDeleteResp{}
	return result, nil
}

func AiCreditLevelConfOnline(ctx context.Context, req *api.AiCreditLevelConfOnlineReq) (*api.AiCreditLevelConfOnlineResp, error) {
	_, _, err := dao.OnlineAiCreditLevelConf(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.AiCreditLevelConfOnlineResp{}
	return result, nil
}
