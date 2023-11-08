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

func AiSysConfigGetAll(ctx context.Context, req *api.AiSysConfigGetAllReq) (*api.AiSysConfigGetAllResp, error) {
	records, totalRaws, err := dao.GetAllAiSysConfig(ctx, req.Page, req.PageNum, req.Order)
	if err != nil {
		return nil, err
	}

	getResp := []*api.AiSysConfigGetResp{}
	for _, v := range records {
		getResp = append(getResp, &api.AiSysConfigGetResp{
			ID: v.ID, ConfigName: v.ConfigName, ConfigKey: v.ConfigKey, ConfigValue: v.ConfigValue, Remark: v.Remark, Deleted: v.Deleted, CreateBy: v.CreateBy, CreateTime: v.CreateTime, UpdateBy: v.UpdateBy, UpdateTime: v.UpdateTime,
		})
	}

	result := &api.AiSysConfigGetAllResp{
		Total: totalRaws,
		List:  getResp,
	}
	return result, nil
}

func AiSysConfigGet(ctx context.Context, req *api.AiSysConfigGetReq) (*api.AiSysConfigGetResp, error) {
	record, err := dao.GetAiSysConfig(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiSysConfigGetResp{
		ID: record.ID, ConfigName: record.ConfigName, ConfigKey: record.ConfigKey, ConfigValue: record.ConfigValue, Remark: record.Remark, Deleted: record.Deleted, CreateBy: record.CreateBy, CreateTime: record.CreateTime, UpdateBy: record.UpdateBy, UpdateTime: record.UpdateTime,
	}

	return result, nil
}

func AiSysConfigUpdate(ctx context.Context, req *api.AiSysConfigUpdateReq) (*api.AiSysConfigUpdateResp, error) {
	updated := &model.AiSysConfig{
		ID: req.ID, ConfigName: req.ConfigName, ConfigKey: req.ConfigKey, ConfigValue: req.ConfigValue, Remark: req.Remark, Deleted: req.Deleted, CreateBy: req.CreateBy, CreateTime: req.CreateTime, UpdateBy: req.UpdateBy, UpdateTime: req.UpdateTime,
	}
	record, _, err := dao.UpdateAiSysConfig(ctx, req.ID, updated)
	if err != nil {
		return nil, err
	}

	result := &api.AiSysConfigUpdateResp{
		ID: record.ID, ConfigName: record.ConfigName, ConfigKey: record.ConfigKey, ConfigValue: record.ConfigValue, Remark: record.Remark, Deleted: record.Deleted, CreateBy: record.CreateBy, CreateTime: record.CreateTime, UpdateBy: record.UpdateBy, UpdateTime: record.UpdateTime,
	}
	return result, nil
}

func AiSysConfigDelete(ctx context.Context, req *api.AiSysConfigDeleteReq) (*api.AiSysConfigDeleteResp, error) {
	_, err := dao.DeleteAiSysConfig(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	result := &api.AiSysConfigDeleteResp{}
	return result, nil
}
