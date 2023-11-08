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

func AiUserVideoRelGetAll(ctx context.Context, req *api.AiUserVideoRelGetAllReq) (*api.AiUserVideoRelGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiUserVideoRel(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiUserVideoRelGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiUserVideoRelGetResp{
			ID: v.ID, UserID: v.UserID, VideoType: v.VideoType, VideoID: v.VideoID, VideoSort: v.VideoSort, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt,
		})
	}

	result := &api.AiUserVideoRelGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiUserVideoRelGet(ctx context.Context, req *api.AiUserVideoRelGetReq) (*api.AiUserVideoRelGetResp, error) {
	record, err := dao.GetAiUserVideoRel(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserVideoRelGetResp{
		ID: record.ID, UserID: record.UserID, VideoType: record.VideoType, VideoID: record.VideoID, VideoSort: record.VideoSort, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}

	return result, nil
}

func AiUserVideoRelUpdate(ctx context.Context, req *api.AiUserVideoRelUpdateReq) (*api.AiUserVideoRelUpdateResp, error) {
	updated := &model.AiUserVideoRel{
		ID: req.ID, UserID: req.UserID, VideoType: req.VideoType, VideoID: req.VideoID, VideoSort: req.VideoSort, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt,
	}
	record, _, err := dao.UpdateAiUserVideoRel(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserVideoRelUpdateResp{
		ID: record.ID, UserID: record.UserID, VideoType: record.VideoType, VideoID: record.VideoID, VideoSort: record.VideoSort, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}
	return result, nil
}

func AiUserVideoRelDelete(ctx context.Context, req *api.AiUserVideoRelDeleteReq) (*api.AiUserVideoRelDeleteResp, error) {
	_, err := dao.DeleteAiUserVideoRel(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserVideoRelDeleteResp{}
	return result, nil
}
