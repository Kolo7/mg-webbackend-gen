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

func AiPlayGuideGetAll(ctx context.Context, req *api.AiPlayGuideGetAllReq) (*api.AiPlayGuideGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiPlayGuide(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiPlayGuideGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiPlayGuideGetResp{
			ID: v.ID, GuideType: v.GuideType, Content: v.Content, GuideJump: v.GuideJump, OnlineStatus: v.OnlineStatus, VideoType: v.VideoType, VideoID: v.VideoID, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt,
		})
	}

	result := &api.AiPlayGuideGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiPlayGuideGet(ctx context.Context, req *api.AiPlayGuideGetReq) (*api.AiPlayGuideGetResp, error) {
	record, err := dao.GetAiPlayGuide(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiPlayGuideGetResp{
		ID: record.ID, GuideType: record.GuideType, Content: record.Content, GuideJump: record.GuideJump, OnlineStatus: record.OnlineStatus, VideoType: record.VideoType, VideoID: record.VideoID, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}

	return result, nil
}

func AiPlayGuideUpdate(ctx context.Context, req *api.AiPlayGuideUpdateReq) (*api.AiPlayGuideUpdateResp, error) {
	updated := &model.AiPlayGuide{
		ID: req.ID, GuideType: req.GuideType, Content: req.Content, GuideJump: req.GuideJump, OnlineStatus: req.OnlineStatus, VideoType: req.VideoType, VideoID: req.VideoID, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt, UpdatedAt: req.UpdatedAt,
	}
	record, _, err := dao.UpdateAiPlayGuide(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiPlayGuideUpdateResp{
		ID: record.ID, GuideType: record.GuideType, Content: record.Content, GuideJump: record.GuideJump, OnlineStatus: record.OnlineStatus, VideoType: record.VideoType, VideoID: record.VideoID, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt, UpdatedAt: record.UpdatedAt,
	}
	return result, nil
}

func AiPlayGuideDelete(ctx context.Context, req *api.AiPlayGuideDeleteReq) (*api.AiPlayGuideDeleteResp, error) {
	_, err := dao.DeleteAiPlayGuide(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiPlayGuideDeleteResp{}
	return result, nil
}

func AiPlayGuideOnline(ctx context.Context, req *api.AiPlayGuideOnlineReq) (*api.AiPlayGuideOnlineResp, error) {
	_, _, err := dao.OnlineAiPlayGuide(ctx, req.ID, req.OnlineStatus)
	if err != nil {
		return nil, err
	}

	result := &api.AiPlayGuideOnlineResp{}
	return result, nil
}
