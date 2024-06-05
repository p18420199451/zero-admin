// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/dictitemservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type DictItemServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedDictItemServiceServer
}

func NewDictItemServiceServer(svcCtx *svc.ServiceContext) *DictItemServiceServer {
	return &DictItemServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加字典数据表
func (s *DictItemServiceServer) AddDictItem(ctx context.Context, in *sysclient.AddDictItemReq) (*sysclient.AddDictItemResp, error) {
	l := dictitemservicelogic.NewAddDictItemLogic(ctx, s.svcCtx)
	return l.AddDictItem(in)
}

// 删除字典数据表
func (s *DictItemServiceServer) DeleteDictItem(ctx context.Context, in *sysclient.DeleteDictItemReq) (*sysclient.DeleteDictItemResp, error) {
	l := dictitemservicelogic.NewDeleteDictItemLogic(ctx, s.svcCtx)
	return l.DeleteDictItem(in)
}

// 更新字典数据表
func (s *DictItemServiceServer) UpdateDictItem(ctx context.Context, in *sysclient.UpdateDictItemReq) (*sysclient.UpdateDictItemResp, error) {
	l := dictitemservicelogic.NewUpdateDictItemLogic(ctx, s.svcCtx)
	return l.UpdateDictItem(in)
}

// 更新字典数据表状态
func (s *DictItemServiceServer) UpdateDictItemStatus(ctx context.Context, in *sysclient.UpdateDictItemStatusReq) (*sysclient.UpdateDictItemStatusResp, error) {
	l := dictitemservicelogic.NewUpdateDictItemStatusLogic(ctx, s.svcCtx)
	return l.UpdateDictItemStatus(in)
}

// 查询字典数据表详情
func (s *DictItemServiceServer) QueryDictItemDetail(ctx context.Context, in *sysclient.QueryDictItemDetailReq) (*sysclient.QueryDictItemDetailResp, error) {
	l := dictitemservicelogic.NewQueryDictItemDetailLogic(ctx, s.svcCtx)
	return l.QueryDictItemDetail(in)
}

// 查询字典数据表列表
func (s *DictItemServiceServer) QueryDictItemList(ctx context.Context, in *sysclient.QueryDictItemListReq) (*sysclient.QueryDictItemListResp, error) {
	l := dictitemservicelogic.NewQueryDictItemListLogic(ctx, s.svcCtx)
	return l.QueryDictItemList(in)
}