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

func newInstNodeTaskFormper(db *gorm.DB, opts ...gen.DOOption) instNodeTaskFormper {
	_instNodeTaskFormper := instNodeTaskFormper{}

	_instNodeTaskFormper.instNodeTaskFormperDo.UseDB(db, opts...)
	_instNodeTaskFormper.instNodeTaskFormperDo.UseModel(&model.InstNodeTaskFormper{})

	tableName := _instNodeTaskFormper.instNodeTaskFormperDo.TableName()
	_instNodeTaskFormper.ALL = field.NewAsterisk(tableName)
	_instNodeTaskFormper.ID = field.NewInt64(tableName, "id")
	_instNodeTaskFormper.InstTaskID = field.NewString(tableName, "inst_task_id")
	_instNodeTaskFormper.NodeTaskID = field.NewString(tableName, "node_task_id")
	_instNodeTaskFormper.NodeID = field.NewString(tableName, "node_id")
	_instNodeTaskFormper.ElemID = field.NewString(tableName, "elemId")
	_instNodeTaskFormper.ElemPID = field.NewString(tableName, "elemPId")
	_instNodeTaskFormper.Per = field.NewInt32(tableName, "per")

	_instNodeTaskFormper.fillFieldMap()

	return _instNodeTaskFormper
}

type instNodeTaskFormper struct {
	instNodeTaskFormperDo instNodeTaskFormperDo

	ALL        field.Asterisk
	ID         field.Int64  // 唯一id
	InstTaskID field.String // 实例任务id
	NodeTaskID field.String // 节点任务id
	NodeID     field.String // 节点id
	ElemID     field.String // 表单元素ID
	ElemPID    field.String // 表单元素父ID
	Per        field.Int32  // 表单权限【可编辑：1；只读：2；隐藏：3】默认只读2

	fieldMap map[string]field.Expr
}

func (i instNodeTaskFormper) Table(newTableName string) *instNodeTaskFormper {
	i.instNodeTaskFormperDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i instNodeTaskFormper) As(alias string) *instNodeTaskFormper {
	i.instNodeTaskFormperDo.DO = *(i.instNodeTaskFormperDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *instNodeTaskFormper) updateTableName(table string) *instNodeTaskFormper {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.InstTaskID = field.NewString(table, "inst_task_id")
	i.NodeTaskID = field.NewString(table, "node_task_id")
	i.NodeID = field.NewString(table, "node_id")
	i.ElemID = field.NewString(table, "elemId")
	i.ElemPID = field.NewString(table, "elemPId")
	i.Per = field.NewInt32(table, "per")

	i.fillFieldMap()

	return i
}

func (i *instNodeTaskFormper) WithContext(ctx context.Context) *instNodeTaskFormperDo {
	return i.instNodeTaskFormperDo.WithContext(ctx)
}

func (i instNodeTaskFormper) TableName() string { return i.instNodeTaskFormperDo.TableName() }

func (i instNodeTaskFormper) Alias() string { return i.instNodeTaskFormperDo.Alias() }

func (i *instNodeTaskFormper) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *instNodeTaskFormper) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 7)
	i.fieldMap["id"] = i.ID
	i.fieldMap["inst_task_id"] = i.InstTaskID
	i.fieldMap["node_task_id"] = i.NodeTaskID
	i.fieldMap["node_id"] = i.NodeID
	i.fieldMap["elemId"] = i.ElemID
	i.fieldMap["elemPId"] = i.ElemPID
	i.fieldMap["per"] = i.Per
}

func (i instNodeTaskFormper) clone(db *gorm.DB) instNodeTaskFormper {
	i.instNodeTaskFormperDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i instNodeTaskFormper) replaceDB(db *gorm.DB) instNodeTaskFormper {
	i.instNodeTaskFormperDo.ReplaceDB(db)
	return i
}

type instNodeTaskFormperDo struct{ gen.DO }

func (i instNodeTaskFormperDo) Debug() *instNodeTaskFormperDo {
	return i.withDO(i.DO.Debug())
}

func (i instNodeTaskFormperDo) WithContext(ctx context.Context) *instNodeTaskFormperDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i instNodeTaskFormperDo) ReadDB() *instNodeTaskFormperDo {
	return i.Clauses(dbresolver.Read)
}

func (i instNodeTaskFormperDo) WriteDB() *instNodeTaskFormperDo {
	return i.Clauses(dbresolver.Write)
}

func (i instNodeTaskFormperDo) Session(config *gorm.Session) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Session(config))
}

func (i instNodeTaskFormperDo) Clauses(conds ...clause.Expression) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i instNodeTaskFormperDo) Returning(value interface{}, columns ...string) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i instNodeTaskFormperDo) Not(conds ...gen.Condition) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i instNodeTaskFormperDo) Or(conds ...gen.Condition) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i instNodeTaskFormperDo) Select(conds ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i instNodeTaskFormperDo) Where(conds ...gen.Condition) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i instNodeTaskFormperDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *instNodeTaskFormperDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i instNodeTaskFormperDo) Order(conds ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i instNodeTaskFormperDo) Distinct(cols ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i instNodeTaskFormperDo) Omit(cols ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i instNodeTaskFormperDo) Join(table schema.Tabler, on ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i instNodeTaskFormperDo) LeftJoin(table schema.Tabler, on ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i instNodeTaskFormperDo) RightJoin(table schema.Tabler, on ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i instNodeTaskFormperDo) Group(cols ...field.Expr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i instNodeTaskFormperDo) Having(conds ...gen.Condition) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i instNodeTaskFormperDo) Limit(limit int) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i instNodeTaskFormperDo) Offset(offset int) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i instNodeTaskFormperDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i instNodeTaskFormperDo) Unscoped() *instNodeTaskFormperDo {
	return i.withDO(i.DO.Unscoped())
}

func (i instNodeTaskFormperDo) Create(values ...*model.InstNodeTaskFormper) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i instNodeTaskFormperDo) CreateInBatches(values []*model.InstNodeTaskFormper, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i instNodeTaskFormperDo) Save(values ...*model.InstNodeTaskFormper) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i instNodeTaskFormperDo) First() (*model.InstNodeTaskFormper, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTaskFormper), nil
	}
}

func (i instNodeTaskFormperDo) Take() (*model.InstNodeTaskFormper, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTaskFormper), nil
	}
}

func (i instNodeTaskFormperDo) Last() (*model.InstNodeTaskFormper, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTaskFormper), nil
	}
}

func (i instNodeTaskFormperDo) Find() ([]*model.InstNodeTaskFormper, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstNodeTaskFormper), err
}

func (i instNodeTaskFormperDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstNodeTaskFormper, err error) {
	buf := make([]*model.InstNodeTaskFormper, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i instNodeTaskFormperDo) FindInBatches(result *[]*model.InstNodeTaskFormper, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i instNodeTaskFormperDo) Attrs(attrs ...field.AssignExpr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i instNodeTaskFormperDo) Assign(attrs ...field.AssignExpr) *instNodeTaskFormperDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i instNodeTaskFormperDo) Joins(fields ...field.RelationField) *instNodeTaskFormperDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i instNodeTaskFormperDo) Preload(fields ...field.RelationField) *instNodeTaskFormperDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i instNodeTaskFormperDo) FirstOrInit() (*model.InstNodeTaskFormper, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTaskFormper), nil
	}
}

func (i instNodeTaskFormperDo) FirstOrCreate() (*model.InstNodeTaskFormper, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTaskFormper), nil
	}
}

func (i instNodeTaskFormperDo) FindByPage(offset int, limit int) (result []*model.InstNodeTaskFormper, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i instNodeTaskFormperDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i instNodeTaskFormperDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i instNodeTaskFormperDo) Delete(models ...*model.InstNodeTaskFormper) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *instNodeTaskFormperDo) withDO(do gen.Dao) *instNodeTaskFormperDo {
	i.DO = *do.(*gen.DO)
	return i
}
