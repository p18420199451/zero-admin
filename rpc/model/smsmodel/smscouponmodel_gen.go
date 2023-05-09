// Code generated by goctl. DO NOT EDIT.

package smsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	smsCouponFieldNames          = builder.RawFieldNames(&SmsCoupon{})
	smsCouponRows                = strings.Join(smsCouponFieldNames, ",")
	smsCouponRowsExpectAutoSet   = strings.Join(stringx.Remove(smsCouponFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	smsCouponRowsWithPlaceHolder = strings.Join(stringx.Remove(smsCouponFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	smsCouponModel interface {
		Insert(ctx context.Context, data *SmsCoupon) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SmsCoupon, error)
		Update(ctx context.Context, data *SmsCoupon) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsCouponModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsCoupon struct {
		Id           int64     `db:"id"`
		Type         int64     `db:"type"`          // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
		Name         string    `db:"name"`          // 名称
		Platform     int64     `db:"platform"`      // 使用平台：0->全部；1->移动；2->PC
		Count        int64     `db:"count"`         // 数量
		Amount       float64   `db:"amount"`        // 金额
		PerLimit     int64     `db:"per_limit"`     // 每人限领张数
		MinPoint     float64   `db:"min_point"`     // 使用门槛；0表示无门槛
		StartTime    time.Time `db:"start_time"`    // 开始时间
		EndTime      time.Time `db:"end_time"`      // 结束时间
		UseType      int64     `db:"use_type"`      // 使用类型：0->全场通用；1->指定分类；2->指定商品
		Note         string    `db:"note"`          // 备注
		PublishCount int64     `db:"publish_count"` // 发行数量
		UseCount     int64     `db:"use_count"`     // 已使用数量
		ReceiveCount int64     `db:"receive_count"` // 领取数量
		EnableTime   time.Time `db:"enable_time"`   // 可以领取的日期
		Code         string    `db:"code"`          // 优惠码
		MemberLevel  int64     `db:"member_level"`  // 可领取的会员类型：0->无限时
	}
)

func newSmsCouponModel(conn sqlx.SqlConn) *defaultSmsCouponModel {
	return &defaultSmsCouponModel{
		conn:  conn,
		table: "`sms_coupon`",
	}
}

func (m *defaultSmsCouponModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSmsCouponModel) FindOne(ctx context.Context, id int64) (*SmsCoupon, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", smsCouponRows, m.table)
	var resp SmsCoupon
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSmsCouponModel) Insert(ctx context.Context, data *SmsCoupon) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, smsCouponRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Type, data.Name, data.Platform, data.Count, data.Amount, data.PerLimit, data.MinPoint, data.StartTime, data.EndTime, data.UseType, data.Note, data.PublishCount, data.UseCount, data.ReceiveCount, data.EnableTime, data.Code, data.MemberLevel)
	return ret, err
}

func (m *defaultSmsCouponModel) Update(ctx context.Context, data *SmsCoupon) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, smsCouponRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Type, data.Name, data.Platform, data.Count, data.Amount, data.PerLimit, data.MinPoint, data.StartTime, data.EndTime, data.UseType, data.Note, data.PublishCount, data.UseCount, data.ReceiveCount, data.EnableTime, data.Code, data.MemberLevel, data.Id)
	return err
}

func (m *defaultSmsCouponModel) tableName() string {
	return m.table
}