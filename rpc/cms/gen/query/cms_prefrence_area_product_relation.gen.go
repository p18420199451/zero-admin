// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/cms/gen/model"
)

func newCmsPrefrenceAreaProductRelation(db *gorm.DB, opts ...gen.DOOption) cmsPrefrenceAreaProductRelation {
	_cmsPrefrenceAreaProductRelation := cmsPrefrenceAreaProductRelation{}

	_cmsPrefrenceAreaProductRelation.cmsPrefrenceAreaProductRelationDo.UseDB(db, opts...)
	_cmsPrefrenceAreaProductRelation.cmsPrefrenceAreaProductRelationDo.UseModel(&model.CmsPrefrenceAreaProductRelation{})

	tableName := _cmsPrefrenceAreaProductRelation.cmsPrefrenceAreaProductRelationDo.TableName()
	_cmsPrefrenceAreaProductRelation.ALL = field.NewAsterisk(tableName)
	_cmsPrefrenceAreaProductRelation.ID = field.NewInt64(tableName, "id")
	_cmsPrefrenceAreaProductRelation.PrefrenceAreaID = field.NewInt64(tableName, "prefrence_area_id")
	_cmsPrefrenceAreaProductRelation.ProductID = field.NewInt64(tableName, "product_id")

	_cmsPrefrenceAreaProductRelation.fillFieldMap()

	return _cmsPrefrenceAreaProductRelation
}

// cmsPrefrenceAreaProductRelation 优选专区和产品关系表
type cmsPrefrenceAreaProductRelation struct {
	cmsPrefrenceAreaProductRelationDo cmsPrefrenceAreaProductRelationDo

	ALL             field.Asterisk
	ID              field.Int64
	PrefrenceAreaID field.Int64
	ProductID       field.Int64

	fieldMap map[string]field.Expr
}

func (c cmsPrefrenceAreaProductRelation) Table(newTableName string) *cmsPrefrenceAreaProductRelation {
	c.cmsPrefrenceAreaProductRelationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsPrefrenceAreaProductRelation) As(alias string) *cmsPrefrenceAreaProductRelation {
	c.cmsPrefrenceAreaProductRelationDo.DO = *(c.cmsPrefrenceAreaProductRelationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsPrefrenceAreaProductRelation) updateTableName(table string) *cmsPrefrenceAreaProductRelation {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.PrefrenceAreaID = field.NewInt64(table, "prefrence_area_id")
	c.ProductID = field.NewInt64(table, "product_id")

	c.fillFieldMap()

	return c
}

func (c *cmsPrefrenceAreaProductRelation) WithContext(ctx context.Context) ICmsPrefrenceAreaProductRelationDo {
	return c.cmsPrefrenceAreaProductRelationDo.WithContext(ctx)
}

func (c cmsPrefrenceAreaProductRelation) TableName() string {
	return c.cmsPrefrenceAreaProductRelationDo.TableName()
}

func (c cmsPrefrenceAreaProductRelation) Alias() string {
	return c.cmsPrefrenceAreaProductRelationDo.Alias()
}

func (c cmsPrefrenceAreaProductRelation) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsPrefrenceAreaProductRelationDo.Columns(cols...)
}

func (c *cmsPrefrenceAreaProductRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsPrefrenceAreaProductRelation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["prefrence_area_id"] = c.PrefrenceAreaID
	c.fieldMap["product_id"] = c.ProductID
}

func (c cmsPrefrenceAreaProductRelation) clone(db *gorm.DB) cmsPrefrenceAreaProductRelation {
	c.cmsPrefrenceAreaProductRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsPrefrenceAreaProductRelation) replaceDB(db *gorm.DB) cmsPrefrenceAreaProductRelation {
	c.cmsPrefrenceAreaProductRelationDo.ReplaceDB(db)
	return c
}

type cmsPrefrenceAreaProductRelationDo struct{ gen.DO }

type ICmsPrefrenceAreaProductRelationDo interface {
	gen.SubQuery
	Debug() ICmsPrefrenceAreaProductRelationDo
	WithContext(ctx context.Context) ICmsPrefrenceAreaProductRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsPrefrenceAreaProductRelationDo
	WriteDB() ICmsPrefrenceAreaProductRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsPrefrenceAreaProductRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsPrefrenceAreaProductRelationDo
	Not(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo
	Or(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo
	Select(conds ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Where(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo
	Order(conds ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Distinct(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Omit(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Join(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Group(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo
	Having(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo
	Limit(limit int) ICmsPrefrenceAreaProductRelationDo
	Offset(offset int) ICmsPrefrenceAreaProductRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsPrefrenceAreaProductRelationDo
	Unscoped() ICmsPrefrenceAreaProductRelationDo
	Create(values ...*model.CmsPrefrenceAreaProductRelation) error
	CreateInBatches(values []*model.CmsPrefrenceAreaProductRelation, batchSize int) error
	Save(values ...*model.CmsPrefrenceAreaProductRelation) error
	First() (*model.CmsPrefrenceAreaProductRelation, error)
	Take() (*model.CmsPrefrenceAreaProductRelation, error)
	Last() (*model.CmsPrefrenceAreaProductRelation, error)
	Find() ([]*model.CmsPrefrenceAreaProductRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsPrefrenceAreaProductRelation, err error)
	FindInBatches(result *[]*model.CmsPrefrenceAreaProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsPrefrenceAreaProductRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsPrefrenceAreaProductRelationDo
	Assign(attrs ...field.AssignExpr) ICmsPrefrenceAreaProductRelationDo
	Joins(fields ...field.RelationField) ICmsPrefrenceAreaProductRelationDo
	Preload(fields ...field.RelationField) ICmsPrefrenceAreaProductRelationDo
	FirstOrInit() (*model.CmsPrefrenceAreaProductRelation, error)
	FirstOrCreate() (*model.CmsPrefrenceAreaProductRelation, error)
	FindByPage(offset int, limit int) (result []*model.CmsPrefrenceAreaProductRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsPrefrenceAreaProductRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsPrefrenceAreaProductRelationDo) Debug() ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsPrefrenceAreaProductRelationDo) WithContext(ctx context.Context) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsPrefrenceAreaProductRelationDo) ReadDB() ICmsPrefrenceAreaProductRelationDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsPrefrenceAreaProductRelationDo) WriteDB() ICmsPrefrenceAreaProductRelationDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsPrefrenceAreaProductRelationDo) Session(config *gorm.Session) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsPrefrenceAreaProductRelationDo) Clauses(conds ...clause.Expression) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Returning(value interface{}, columns ...string) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsPrefrenceAreaProductRelationDo) Not(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Or(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Select(conds ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Where(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Order(conds ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Distinct(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsPrefrenceAreaProductRelationDo) Omit(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsPrefrenceAreaProductRelationDo) Join(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsPrefrenceAreaProductRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsPrefrenceAreaProductRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsPrefrenceAreaProductRelationDo) Group(cols ...field.Expr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsPrefrenceAreaProductRelationDo) Having(conds ...gen.Condition) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsPrefrenceAreaProductRelationDo) Limit(limit int) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsPrefrenceAreaProductRelationDo) Offset(offset int) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsPrefrenceAreaProductRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsPrefrenceAreaProductRelationDo) Unscoped() ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsPrefrenceAreaProductRelationDo) Create(values ...*model.CmsPrefrenceAreaProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsPrefrenceAreaProductRelationDo) CreateInBatches(values []*model.CmsPrefrenceAreaProductRelation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsPrefrenceAreaProductRelationDo) Save(values ...*model.CmsPrefrenceAreaProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsPrefrenceAreaProductRelationDo) First() (*model.CmsPrefrenceAreaProductRelation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPrefrenceAreaProductRelation), nil
	}
}

func (c cmsPrefrenceAreaProductRelationDo) Take() (*model.CmsPrefrenceAreaProductRelation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPrefrenceAreaProductRelation), nil
	}
}

func (c cmsPrefrenceAreaProductRelationDo) Last() (*model.CmsPrefrenceAreaProductRelation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPrefrenceAreaProductRelation), nil
	}
}

func (c cmsPrefrenceAreaProductRelationDo) Find() ([]*model.CmsPrefrenceAreaProductRelation, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsPrefrenceAreaProductRelation), err
}

func (c cmsPrefrenceAreaProductRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsPrefrenceAreaProductRelation, err error) {
	buf := make([]*model.CmsPrefrenceAreaProductRelation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsPrefrenceAreaProductRelationDo) FindInBatches(result *[]*model.CmsPrefrenceAreaProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsPrefrenceAreaProductRelationDo) Attrs(attrs ...field.AssignExpr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsPrefrenceAreaProductRelationDo) Assign(attrs ...field.AssignExpr) ICmsPrefrenceAreaProductRelationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsPrefrenceAreaProductRelationDo) Joins(fields ...field.RelationField) ICmsPrefrenceAreaProductRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsPrefrenceAreaProductRelationDo) Preload(fields ...field.RelationField) ICmsPrefrenceAreaProductRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsPrefrenceAreaProductRelationDo) FirstOrInit() (*model.CmsPrefrenceAreaProductRelation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPrefrenceAreaProductRelation), nil
	}
}

func (c cmsPrefrenceAreaProductRelationDo) FirstOrCreate() (*model.CmsPrefrenceAreaProductRelation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPrefrenceAreaProductRelation), nil
	}
}

func (c cmsPrefrenceAreaProductRelationDo) FindByPage(offset int, limit int) (result []*model.CmsPrefrenceAreaProductRelation, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cmsPrefrenceAreaProductRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsPrefrenceAreaProductRelationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsPrefrenceAreaProductRelationDo) Delete(models ...*model.CmsPrefrenceAreaProductRelation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsPrefrenceAreaProductRelationDo) withDO(do gen.Dao) *cmsPrefrenceAreaProductRelationDo {
	c.DO = *do.(*gen.DO)
	return c
}