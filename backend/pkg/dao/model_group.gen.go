// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/wegoteam/weflow/pkg/model"
)

func newModelGroup(db *gorm.DB, opts ...gen.DOOption) modelGroup {
	_modelGroup := modelGroup{}

	_modelGroup.modelGroupDo.UseDB(db, opts...)
	_modelGroup.modelGroupDo.UseModel(&model.ModelGroup{})

	tableName := _modelGroup.modelGroupDo.TableName()
	_modelGroup.ALL = field.NewAsterisk(tableName)
	_modelGroup.ID = field.NewInt64(tableName, "id")
	_modelGroup.GroupID = field.NewString(tableName, "group_id")
	_modelGroup.GroupName = field.NewString(tableName, "group_name")
	_modelGroup.Remark = field.NewString(tableName, "remark")
	_modelGroup.CreateUser = field.NewString(tableName, "create_user")
	_modelGroup.UpdateUser = field.NewString(tableName, "update_user")
	_modelGroup.CreateTime = field.NewTime(tableName, "create_time")
	_modelGroup.UpdateTime = field.NewTime(tableName, "update_time")

	_modelGroup.fillFieldMap()

	return _modelGroup
}

type modelGroup struct {
	modelGroupDo modelGroupDo

	ALL        field.Asterisk
	ID         field.Int64  // 唯一id
	GroupID    field.String // 组id
	GroupName  field.String // 组名称
	Remark     field.String // 描述
	CreateUser field.String // 创建人
	UpdateUser field.String // 更新人
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (m modelGroup) Table(newTableName string) *modelGroup {
	m.modelGroupDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m modelGroup) As(alias string) *modelGroup {
	m.modelGroupDo.DO = *(m.modelGroupDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *modelGroup) updateTableName(table string) *modelGroup {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.GroupID = field.NewString(table, "group_id")
	m.GroupName = field.NewString(table, "group_name")
	m.Remark = field.NewString(table, "remark")
	m.CreateUser = field.NewString(table, "create_user")
	m.UpdateUser = field.NewString(table, "update_user")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")

	m.fillFieldMap()

	return m
}

func (m *modelGroup) WithContext(ctx context.Context) *modelGroupDo {
	return m.modelGroupDo.WithContext(ctx)
}

func (m modelGroup) TableName() string { return m.modelGroupDo.TableName() }

func (m modelGroup) Alias() string { return m.modelGroupDo.Alias() }

func (m *modelGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *modelGroup) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.ID
	m.fieldMap["group_id"] = m.GroupID
	m.fieldMap["group_name"] = m.GroupName
	m.fieldMap["remark"] = m.Remark
	m.fieldMap["create_user"] = m.CreateUser
	m.fieldMap["update_user"] = m.UpdateUser
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
}

func (m modelGroup) clone(db *gorm.DB) modelGroup {
	m.modelGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m modelGroup) replaceDB(db *gorm.DB) modelGroup {
	m.modelGroupDo.ReplaceDB(db)
	return m
}

type modelGroupDo struct{ gen.DO }

func (m modelGroupDo) Debug() *modelGroupDo {
	return m.withDO(m.DO.Debug())
}

func (m modelGroupDo) WithContext(ctx context.Context) *modelGroupDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m modelGroupDo) ReadDB() *modelGroupDo {
	return m.Clauses(dbresolver.Read)
}

func (m modelGroupDo) WriteDB() *modelGroupDo {
	return m.Clauses(dbresolver.Write)
}

func (m modelGroupDo) Session(config *gorm.Session) *modelGroupDo {
	return m.withDO(m.DO.Session(config))
}

func (m modelGroupDo) Clauses(conds ...clause.Expression) *modelGroupDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m modelGroupDo) Returning(value interface{}, columns ...string) *modelGroupDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m modelGroupDo) Not(conds ...gen.Condition) *modelGroupDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m modelGroupDo) Or(conds ...gen.Condition) *modelGroupDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m modelGroupDo) Select(conds ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m modelGroupDo) Where(conds ...gen.Condition) *modelGroupDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m modelGroupDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *modelGroupDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m modelGroupDo) Order(conds ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m modelGroupDo) Distinct(cols ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m modelGroupDo) Omit(cols ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m modelGroupDo) Join(table schema.Tabler, on ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m modelGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m modelGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m modelGroupDo) Group(cols ...field.Expr) *modelGroupDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m modelGroupDo) Having(conds ...gen.Condition) *modelGroupDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m modelGroupDo) Limit(limit int) *modelGroupDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m modelGroupDo) Offset(offset int) *modelGroupDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m modelGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *modelGroupDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m modelGroupDo) Unscoped() *modelGroupDo {
	return m.withDO(m.DO.Unscoped())
}

func (m modelGroupDo) Create(values ...*model.ModelGroup) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m modelGroupDo) CreateInBatches(values []*model.ModelGroup, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m modelGroupDo) Save(values ...*model.ModelGroup) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m modelGroupDo) First() (*model.ModelGroup, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelGroup), nil
	}
}

func (m modelGroupDo) Take() (*model.ModelGroup, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelGroup), nil
	}
}

func (m modelGroupDo) Last() (*model.ModelGroup, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelGroup), nil
	}
}

func (m modelGroupDo) Find() ([]*model.ModelGroup, error) {
	result, err := m.DO.Find()
	return result.([]*model.ModelGroup), err
}

func (m modelGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ModelGroup, err error) {
	buf := make([]*model.ModelGroup, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m modelGroupDo) FindInBatches(result *[]*model.ModelGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m modelGroupDo) Attrs(attrs ...field.AssignExpr) *modelGroupDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m modelGroupDo) Assign(attrs ...field.AssignExpr) *modelGroupDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m modelGroupDo) Joins(fields ...field.RelationField) *modelGroupDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m modelGroupDo) Preload(fields ...field.RelationField) *modelGroupDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m modelGroupDo) FirstOrInit() (*model.ModelGroup, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelGroup), nil
	}
}

func (m modelGroupDo) FirstOrCreate() (*model.ModelGroup, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModelGroup), nil
	}
}

func (m modelGroupDo) FindByPage(offset int, limit int) (result []*model.ModelGroup, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m modelGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m modelGroupDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m modelGroupDo) Delete(models ...*model.ModelGroup) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *modelGroupDo) withDO(do gen.Dao) *modelGroupDo {
	m.DO = *do.(*gen.DO)
	return m
}
