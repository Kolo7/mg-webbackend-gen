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

func ChatKeywordMaterialGetAll(ctx context.Context, req *api.ChatKeywordMaterialGetAllReq) (*api.ChatKeywordMaterialGetAllResp, error) {
	records, totalRaws, err := dao.GetAllChatKeywordMaterial(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.ChatKeywordMaterialGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.ChatKeywordMaterialGetResp{
			ID: v.ID, Uuids: v.Uuids, MatchKeywords: v.MatchKeywords, MaterialType: v.MaterialType, MaterialText: v.MaterialText, TimeLength: v.TimeLength, AudioURL: v.AudioURL, ImageHeight: v.ImageHeight, ImageWidth: v.ImageWidth, ImageURL: v.ImageURL, OnlineStatus: v.OnlineStatus, CreateBy: v.CreateBy, UpdateBy: v.UpdateBy, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.ChatKeywordMaterialGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func ChatKeywordMaterialGet(ctx context.Context, req *api.ChatKeywordMaterialGetReq) (*api.ChatKeywordMaterialGetResp, error) {
	record, err := dao.GetChatKeywordMaterial(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.ChatKeywordMaterialGetResp{
		ID: record.ID, Uuids: record.Uuids, MatchKeywords: record.MatchKeywords, MaterialType: record.MaterialType, MaterialText: record.MaterialText, TimeLength: record.TimeLength, AudioURL: record.AudioURL, ImageHeight: record.ImageHeight, ImageWidth: record.ImageWidth, ImageURL: record.ImageURL, OnlineStatus: record.OnlineStatus, CreateBy: record.CreateBy, UpdateBy: record.UpdateBy, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func ChatKeywordMaterialUpdate(ctx context.Context, req *api.ChatKeywordMaterialUpdateReq) (*api.ChatKeywordMaterialUpdateResp, error) {
	updated := &model.ChatKeywordMaterial{
		ID: req.ID, Uuids: req.Uuids, MatchKeywords: req.MatchKeywords, MaterialType: req.MaterialType, MaterialText: req.MaterialText, TimeLength: req.TimeLength, AudioURL: req.AudioURL, ImageHeight: req.ImageHeight, ImageWidth: req.ImageWidth, ImageURL: req.ImageURL, OnlineStatus: req.OnlineStatus, CreateBy: req.CreateBy, UpdateBy: req.UpdateBy, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateChatKeywordMaterial(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.ChatKeywordMaterialUpdateResp{
		ID: record.ID, Uuids: record.Uuids, MatchKeywords: record.MatchKeywords, MaterialType: record.MaterialType, MaterialText: record.MaterialText, TimeLength: record.TimeLength, AudioURL: record.AudioURL, ImageHeight: record.ImageHeight, ImageWidth: record.ImageWidth, ImageURL: record.ImageURL, OnlineStatus: record.OnlineStatus, CreateBy: record.CreateBy, UpdateBy: record.UpdateBy, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func ChatKeywordMaterialDelete(ctx context.Context, req *api.ChatKeywordMaterialDeleteReq) (*api.ChatKeywordMaterialDeleteResp, error) {
	_, err := dao.DeleteChatKeywordMaterial(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.ChatKeywordMaterialDeleteResp{}
	return result, nil
}

func ChatKeywordMaterialOnline(ctx context.Context, req *api.ChatKeywordMaterialOnlineReq) (*api.ChatKeywordMaterialOnlineResp, error) {
	_, _, err := dao.OnlineChatKeywordMaterial(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.ChatKeywordMaterialOnlineResp{}
	return result, nil
}
