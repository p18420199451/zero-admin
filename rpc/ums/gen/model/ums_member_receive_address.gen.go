// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsMemberReceiveAddress = "ums_member_receive_address"

// UmsMemberReceiveAddress 会员收货地址表
type UmsMemberReceiveAddress struct {
	ID            int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID      int64      `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                               // 会员id
	MemberName    string     `gorm:"column:member_name;not null;comment:收货人名称" json:"member_name"`                          // 收货人名称
	PhoneNumber   string     `gorm:"column:phone_number;not null;comment:收货人电话" json:"phone_number"`                        // 收货人电话
	DefaultStatus int32      `gorm:"column:default_status;not null;comment:是否为默认" json:"default_status"`                    // 是否为默认
	PostCode      string     `gorm:"column:post_code;not null;comment:邮政编码" json:"post_code"`                               // 邮政编码
	Province      string     `gorm:"column:province;not null;comment:省份/直辖市" json:"province"`                               // 省份/直辖市
	City          string     `gorm:"column:city;not null;comment:城市" json:"city"`                                           // 城市
	Region        string     `gorm:"column:region;not null;comment:区" json:"region"`                                        // 区
	DetailAddress string     `gorm:"column:detail_address;not null;comment:详细地址(街道)" json:"detail_address"`                 // 详细地址(街道)
	CreateTime    time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime    *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName UmsMemberReceiveAddress's table name
func (*UmsMemberReceiveAddress) TableName() string {
	return TableNameUmsMemberReceiveAddress
}