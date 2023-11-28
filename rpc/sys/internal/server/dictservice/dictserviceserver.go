// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-admin/rpc/sys/internal/logic/dictservice"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
)

type DictServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedDictServiceServer
}

func NewDictServiceServer(svcCtx *svc.ServiceContext) *DictServiceServer {
	return &DictServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *DictServiceServer) DictAdd(ctx context.Context, in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	l := dictservicelogic.NewDictAddLogic(ctx, s.svcCtx)
	return l.DictAdd(in)
}

func (s *DictServiceServer) DictList(ctx context.Context, in *sysclient.DictListReq) (*sysclient.DictListResp, error) {
	l := dictservicelogic.NewDictListLogic(ctx, s.svcCtx)
	return l.DictList(in)
}

func (s *DictServiceServer) DictUpdate(ctx context.Context, in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	l := dictservicelogic.NewDictUpdateLogic(ctx, s.svcCtx)
	return l.DictUpdate(in)
}

func (s *DictServiceServer) DictDelete(ctx context.Context, in *sysclient.DictDeleteReq) (*sysclient.DictDeleteResp, error) {
	l := dictservicelogic.NewDictDeleteLogic(ctx, s.svcCtx)
	return l.DictDelete(in)
}