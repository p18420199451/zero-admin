package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleLogic 更新用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:38
*/
type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserRole 更新用户与角色的关联(角色id为1的是系统预留超级管理员角色,不用关联)
func (l *UpdateUserRoleLogic) UpdateUserRole(in *sysclient.UpdateUserRoleReq) (*sysclient.UpdateUserRoleResp, error) {
	//判断是否为超级管理员
	//id为1的是系统预留超级管理员角色,不用关联
	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysUserRole
		userId := in.UserId
		count, err := q.WithContext(l.ctx).Where(q.UserID.Eq(userId), q.RoleID.Eq(1)).Count()

		if err != nil {
			return err
		}

		if count != 0 {
			return nil
		}

		if _, err = q.WithContext(l.ctx).Where(q.RoleID.Eq(userId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var userRoles []*model.SysUserRole
		for _, roleId := range in.RoleIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: roleId,
				UserID: userId,
			})
		}

		if err = q.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
			logc.Errorf(l.ctx, "添加用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.UpdateUserRoleResp{}, nil
}
