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

func newProcessDefNodeFormper(db *gorm.DB, opts ...gen.DOOption) processDefNodeFormper {
	_processDefNodeFormper := processDefNodeFormper{}

	_processDefNodeFormper.processDefNodeFormperDo.UseDB(db, opts...)
	_processDefNodeFormper.processDefNodeFormperDo.UseModel(&model.ProcessDefNodeFormper{})

	tableName := _processDefNodeFormper.processDefNodeFormperDo.TableName()
	_processDefNodeFormper.ALL = field.NewAsterisk(tableName)
	_processDefNodeFormper.ID = field.NewInt64(tableName, "id")
	_processDefNodeFormper.ProcessDefID = field.NewString(tableName, "process_def_id")
	_processDefNodeFormper.NodeID = field.NewString(tableName, "node_id")
	_processDefNodeFormper.ElemID = field.NewString(tableName, "elemId")
	_processDefNodeFormper.ElemPID = field.NewString(tableName, "elemPId")
	_processDefNodeFormper.Per = field.NewInt32(tableName, "per")

	_processDefNodeFormper.fillFieldMap()

	return _processDefNodeFormper
}

type processDefNodeFormper struct {
	processDefNodeFormperDo processDefNodeFormperDo

	ALL          field.Asterisk
	ID           field.Int64  // 唯一id
	ProcessDefID field.String // 流程定义id
	NodeID       field.String // 节点id
	ElemID       field.String // 处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等
	ElemPID      field.String // 处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等
	Per          field.Int32  // 处理人顺序;正序排序

	fieldMap map[string]field.Expr
}

func (p processDefNodeFormper) Table(newTableName string) *processDefNodeFormper {
	p.processDefNodeFormperDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p processDefNodeFormper) As(alias string) *processDefNodeFormper {
	p.processDefNodeFormperDo.DO = *(p.processDefNodeFormperDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *processDefNodeFormper) updateTableName(table string) *processDefNodeFormper {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProcessDefID = field.NewString(table, "process_def_id")
	p.NodeID = field.NewString(table, "node_id")
	p.ElemID = field.NewString(table, "elemId")
	p.ElemPID = field.NewString(table, "elemPId")
	p.Per = field.NewInt32(table, "per")

	p.fillFieldMap()

	return p
}

func (p *processDefNodeFormper) WithContext(ctx context.Context) *processDefNodeFormperDo {
	return p.processDefNodeFormperDo.WithContext(ctx)
}

func (p processDefNodeFormper) TableName() string { return p.processDefNodeFormperDo.TableName() }

func (p processDefNodeFormper) Alias() string { return p.processDefNodeFormperDo.Alias() }

func (p *processDefNodeFormper) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *processDefNodeFormper) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["process_def_id"] = p.ProcessDefID
	p.fieldMap["node_id"] = p.NodeID
	p.fieldMap["elemId"] = p.ElemID
	p.fieldMap["elemPId"] = p.ElemPID
	p.fieldMap["per"] = p.Per
}

func (p processDefNodeFormper) clone(db *gorm.DB) processDefNodeFormper {
	p.processDefNodeFormperDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p processDefNodeFormper) replaceDB(db *gorm.DB) processDefNodeFormper {
	p.processDefNodeFormperDo.ReplaceDB(db)
	return p
}

type processDefNodeFormperDo struct{ gen.DO }

func (p processDefNodeFormperDo) Debug() *processDefNodeFormperDo {
	return p.withDO(p.DO.Debug())
}

func (p processDefNodeFormperDo) WithContext(ctx context.Context) *processDefNodeFormperDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p processDefNodeFormperDo) ReadDB() *processDefNodeFormperDo {
	return p.Clauses(dbresolver.Read)
}

func (p processDefNodeFormperDo) WriteDB() *processDefNodeFormperDo {
	return p.Clauses(dbresolver.Write)
}

func (p processDefNodeFormperDo) Session(config *gorm.Session) *processDefNodeFormperDo {
	return p.withDO(p.DO.Session(config))
}

func (p processDefNodeFormperDo) Clauses(conds ...clause.Expression) *processDefNodeFormperDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p processDefNodeFormperDo) Returning(value interface{}, columns ...string) *processDefNodeFormperDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p processDefNodeFormperDo) Not(conds ...gen.Condition) *processDefNodeFormperDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p processDefNodeFormperDo) Or(conds ...gen.Condition) *processDefNodeFormperDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p processDefNodeFormperDo) Select(conds ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p processDefNodeFormperDo) Where(conds ...gen.Condition) *processDefNodeFormperDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p processDefNodeFormperDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *processDefNodeFormperDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p processDefNodeFormperDo) Order(conds ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p processDefNodeFormperDo) Distinct(cols ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p processDefNodeFormperDo) Omit(cols ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p processDefNodeFormperDo) Join(table schema.Tabler, on ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p processDefNodeFormperDo) LeftJoin(table schema.Tabler, on ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p processDefNodeFormperDo) RightJoin(table schema.Tabler, on ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p processDefNodeFormperDo) Group(cols ...field.Expr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p processDefNodeFormperDo) Having(conds ...gen.Condition) *processDefNodeFormperDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p processDefNodeFormperDo) Limit(limit int) *processDefNodeFormperDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p processDefNodeFormperDo) Offset(offset int) *processDefNodeFormperDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p processDefNodeFormperDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *processDefNodeFormperDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p processDefNodeFormperDo) Unscoped() *processDefNodeFormperDo {
	return p.withDO(p.DO.Unscoped())
}

func (p processDefNodeFormperDo) Create(values ...*model.ProcessDefNodeFormper) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p processDefNodeFormperDo) CreateInBatches(values []*model.ProcessDefNodeFormper, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p processDefNodeFormperDo) Save(values ...*model.ProcessDefNodeFormper) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p processDefNodeFormperDo) First() (*model.ProcessDefNodeFormper, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeFormper), nil
	}
}

func (p processDefNodeFormperDo) Take() (*model.ProcessDefNodeFormper, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeFormper), nil
	}
}

func (p processDefNodeFormperDo) Last() (*model.ProcessDefNodeFormper, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeFormper), nil
	}
}

func (p processDefNodeFormperDo) Find() ([]*model.ProcessDefNodeFormper, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProcessDefNodeFormper), err
}

func (p processDefNodeFormperDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProcessDefNodeFormper, err error) {
	buf := make([]*model.ProcessDefNodeFormper, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p processDefNodeFormperDo) FindInBatches(result *[]*model.ProcessDefNodeFormper, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p processDefNodeFormperDo) Attrs(attrs ...field.AssignExpr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p processDefNodeFormperDo) Assign(attrs ...field.AssignExpr) *processDefNodeFormperDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p processDefNodeFormperDo) Joins(fields ...field.RelationField) *processDefNodeFormperDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p processDefNodeFormperDo) Preload(fields ...field.RelationField) *processDefNodeFormperDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p processDefNodeFormperDo) FirstOrInit() (*model.ProcessDefNodeFormper, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeFormper), nil
	}
}

func (p processDefNodeFormperDo) FirstOrCreate() (*model.ProcessDefNodeFormper, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeFormper), nil
	}
}

func (p processDefNodeFormperDo) FindByPage(offset int, limit int) (result []*model.ProcessDefNodeFormper, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p processDefNodeFormperDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p processDefNodeFormperDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p processDefNodeFormperDo) Delete(models ...*model.ProcessDefNodeFormper) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *processDefNodeFormperDo) withDO(do gen.Dao) *processDefNodeFormperDo {
	p.DO = *do.(*gen.DO)
	return p
}