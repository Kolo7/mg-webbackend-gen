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

func AiUserSocialRelGetAll(ctx context.Context, req *api.AiUserSocialRelGetAllReq) (*api.AiUserSocialRelGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiUserSocialRel(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiUserSocialRelGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiUserSocialRelGetResp{
			ID: v.ID, UserID: v.UserID, BindUserID: v.BindUserID, SocialName: v.SocialName, SocialSort: v.SocialSort, Creater: v.Creater, Upadter: v.Upadter, CreatedAt: v.CreatedAt,
		})
	}

	result := &api.AiUserSocialRelGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiUserSocialRelGet(ctx context.Context, req *api.AiUserSocialRelGetReq) (*api.AiUserSocialRelGetResp, error) {
	record, err := dao.GetAiUserSocialRel(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserSocialRelGetResp{
		ID: record.ID, UserID: record.UserID, BindUserID: record.BindUserID, SocialName: record.SocialName, SocialSort: record.SocialSort, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}

	return result, nil
}

func AiUserSocialRelUpdate(ctx context.Context, req *api.AiUserSocialRelUpdateReq) (*api.AiUserSocialRelUpdateResp, error) {
	updated := &model.AiUserSocialRel{
		ID: req.ID, UserID: req.UserID, BindUserID: req.BindUserID, SocialName: req.SocialName, SocialSort: req.SocialSort, Creater: req.Creater, Upadter: req.Upadter, CreatedAt: req.CreatedAt,
	}
	record, _, err := dao.UpdateAiUserSocialRel(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserSocialRelUpdateResp{
		ID: record.ID, UserID: record.UserID, BindUserID: record.BindUserID, SocialName: record.SocialName, SocialSort: record.SocialSort, Creater: record.Creater, Upadter: record.Upadter, CreatedAt: record.CreatedAt,
	}
	return result, nil
}

func AiUserSocialRelDelete(ctx context.Context, req *api.AiUserSocialRelDeleteReq) (*api.AiUserSocialRelDeleteResp, error) {
	_, err := dao.DeleteAiUserSocialRel(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiUserSocialRelDeleteResp{}
	return result, nil
}
