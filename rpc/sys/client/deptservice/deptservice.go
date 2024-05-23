// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package deptservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeptAddReq             = sysclient.DeptAddReq
	DeptAddResp            = sysclient.DeptAddResp
	DeptDeleteReq          = sysclient.DeptDeleteReq
	DeptDeleteResp         = sysclient.DeptDeleteResp
	DeptListData           = sysclient.DeptListData
	DeptListReq            = sysclient.DeptListReq
	DeptListResp           = sysclient.DeptListResp
	DeptUpdateReq          = sysclient.DeptUpdateReq
	DeptUpdateResp         = sysclient.DeptUpdateResp
	DictAddReq             = sysclient.DictAddReq
	DictAddResp            = sysclient.DictAddResp
	DictDeleteReq          = sysclient.DictDeleteReq
	DictDeleteResp         = sysclient.DictDeleteResp
	DictListData           = sysclient.DictListData
	DictListReq            = sysclient.DictListReq
	DictListResp           = sysclient.DictListResp
	DictUpdateReq          = sysclient.DictUpdateReq
	DictUpdateResp         = sysclient.DictUpdateResp
	InfoReq                = sysclient.InfoReq
	InfoResp               = sysclient.InfoResp
	JobAddReq              = sysclient.JobAddReq
	JobAddResp             = sysclient.JobAddResp
	JobDeleteReq           = sysclient.JobDeleteReq
	JobDeleteResp          = sysclient.JobDeleteResp
	JobListData            = sysclient.JobListData
	JobListReq             = sysclient.JobListReq
	JobListResp            = sysclient.JobListResp
	JobUpdateReq           = sysclient.JobUpdateReq
	JobUpdateResp          = sysclient.JobUpdateResp
	LoginLogAddReq         = sysclient.LoginLogAddReq
	LoginLogAddResp        = sysclient.LoginLogAddResp
	LoginLogDeleteReq      = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp     = sysclient.LoginLogDeleteResp
	LoginLogListData       = sysclient.LoginLogListData
	LoginLogListReq        = sysclient.LoginLogListReq
	LoginLogListResp       = sysclient.LoginLogListResp
	LoginReq               = sysclient.LoginReq
	LoginResp              = sysclient.LoginResp
	MenuAddReq             = sysclient.MenuAddReq
	MenuAddResp            = sysclient.MenuAddResp
	MenuDeleteReq          = sysclient.MenuDeleteReq
	MenuDeleteResp         = sysclient.MenuDeleteResp
	MenuListData           = sysclient.MenuListData
	MenuListReq            = sysclient.MenuListReq
	MenuListResp           = sysclient.MenuListResp
	MenuListTree           = sysclient.MenuListTree
	MenuUpdateReq          = sysclient.MenuUpdateReq
	MenuUpdateResp         = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq   = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp  = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq       = sysclient.ReSetPasswordReq
	ReSetPasswordResp      = sysclient.ReSetPasswordResp
	RoleAddReq             = sysclient.RoleAddReq
	RoleAddResp            = sysclient.RoleAddResp
	RoleDeleteReq          = sysclient.RoleDeleteReq
	RoleDeleteResp         = sysclient.RoleDeleteResp
	RoleListData           = sysclient.RoleListData
	RoleListReq            = sysclient.RoleListReq
	RoleListResp           = sysclient.RoleListResp
	RoleUpdateReq          = sysclient.RoleUpdateReq
	RoleUpdateResp         = sysclient.RoleUpdateResp
	StatisticsLoginLogReq  = sysclient.StatisticsLoginLogReq
	StatisticsLoginLogResp = sysclient.StatisticsLoginLogResp
	SysLogAddReq           = sysclient.SysLogAddReq
	SysLogAddResp          = sysclient.SysLogAddResp
	SysLogDeleteReq        = sysclient.SysLogDeleteReq
	SysLogDeleteResp       = sysclient.SysLogDeleteResp
	SysLogListData         = sysclient.SysLogListData
	SysLogListReq          = sysclient.SysLogListReq
	SysLogListResp         = sysclient.SysLogListResp
	UpdateMenuRoleReq      = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp     = sysclient.UpdateMenuRoleResp
	UpdateUserRoleReq      = sysclient.UpdateUserRoleReq
	UpdateUserRoleResp     = sysclient.UpdateUserRoleResp
	UserAddReq             = sysclient.UserAddReq
	UserAddResp            = sysclient.UserAddResp
	UserDeleteReq          = sysclient.UserDeleteReq
	UserDeleteResp         = sysclient.UserDeleteResp
	UserListData           = sysclient.UserListData
	UserListReq            = sysclient.UserListReq
	UserListResp           = sysclient.UserListResp
	UserStatusReq          = sysclient.UserStatusReq
	UserStatusResp         = sysclient.UserStatusResp
	UserUpdateReq          = sysclient.UserUpdateReq
	UserUpdateResp         = sysclient.UserUpdateResp

	DeptService interface {
		DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error)
		DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
		DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error)
		DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error)
	}

	defaultDeptService struct {
		cli zrpc.Client
	}
)

func NewDeptService(cli zrpc.Client) DeptService {
	return &defaultDeptService{
		cli: cli,
	}
}

func (m *defaultDeptService) DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.DeptAdd(ctx, in, opts...)
}

func (m *defaultDeptService) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.DeptList(ctx, in, opts...)
}

func (m *defaultDeptService) DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in, opts...)
}

func (m *defaultDeptService) DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.DeptDelete(ctx, in, opts...)
}
