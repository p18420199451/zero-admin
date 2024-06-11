// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/helpservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type HelpServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedHelpServiceServer
}

func NewHelpServiceServer(svcCtx *svc.ServiceContext) *HelpServiceServer {
	return &HelpServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加帮助表
func (s *HelpServiceServer) AddHelp(ctx context.Context, in *cmsclient.AddHelpReq) (*cmsclient.AddHelpResp, error) {
	l := helpservicelogic.NewAddHelpLogic(ctx, s.svcCtx)
	return l.AddHelp(in)
}

// 删除帮助表
func (s *HelpServiceServer) DeleteHelp(ctx context.Context, in *cmsclient.DeleteHelpReq) (*cmsclient.DeleteHelpResp, error) {
	l := helpservicelogic.NewDeleteHelpLogic(ctx, s.svcCtx)
	return l.DeleteHelp(in)
}

// 更新帮助表
func (s *HelpServiceServer) UpdateHelp(ctx context.Context, in *cmsclient.UpdateHelpReq) (*cmsclient.UpdateHelpResp, error) {
	l := helpservicelogic.NewUpdateHelpLogic(ctx, s.svcCtx)
	return l.UpdateHelp(in)
}

// 更新帮助表状态
func (s *HelpServiceServer) UpdateHelpStatus(ctx context.Context, in *cmsclient.UpdateHelpStatusReq) (*cmsclient.UpdateHelpStatusResp, error) {
	l := helpservicelogic.NewUpdateHelpStatusLogic(ctx, s.svcCtx)
	return l.UpdateHelpStatus(in)
}

// 查询帮助表详情
func (s *HelpServiceServer) QueryHelpDetail(ctx context.Context, in *cmsclient.QueryHelpDetailReq) (*cmsclient.QueryHelpDetailResp, error) {
	l := helpservicelogic.NewQueryHelpDetailLogic(ctx, s.svcCtx)
	return l.QueryHelpDetail(in)
}

// 查询帮助表列表
func (s *HelpServiceServer) QueryHelpList(ctx context.Context, in *cmsclient.QueryHelpListReq) (*cmsclient.QueryHelpListResp, error) {
	l := helpservicelogic.NewQueryHelpListLogic(ctx, s.svcCtx)
	return l.QueryHelpList(in)
}