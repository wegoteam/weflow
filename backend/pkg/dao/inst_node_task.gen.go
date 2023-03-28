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

	"wego2023/weflow/pkg/model"
)

func newInstNodeTask(db *gorm.DB, opts ...gen.DOOption) instNodeTask {
	_instNodeTask := instNodeTask{}

	_instNodeTask.instNodeTaskDo.UseDB(db, opts...)
	_instNodeTask.instNodeTaskDo.UseModel(&model.InstNodeTask{})

	tableName := _instNodeTask.instNodeTaskDo.TableName()
	_instNodeTask.ALL = field.NewAsterisk(tableName)
	_instNodeTask.ID = field.NewInt64(tableName, "id")
	_instNodeTask.InstTaskID = field.NewString(tableName, "inst_task_id")
	_instNodeTask.NodeTaskID = field.NewString(tableName, "node_task_id")
	_instNodeTask.NodeID = field.NewString(tableName, "node_id")
	_instNodeTask.ParentID = field.NewString(tableName, "parent_id")
	_instNodeTask.NodeType = field.NewInt32(tableName, "node_type")
	_instNodeTask.NodeName = field.NewString(tableName, "node_name")
	_instNodeTask.ForwardMode = field.NewInt32(tableName, "forward_mode")
	_instNodeTask.CompleteConn = field.NewInt32(tableName, "complete_conn")
	_instNodeTask.PermissionMode = field.NewInt32(tableName, "permission_mode")
	_instNodeTask.AllowAdd = field.NewInt32(tableName, "allow_add")
	_instNodeTask.ProcessMode = field.NewInt32(tableName, "process_mode")
	_instNodeTask.TimeLimit = field.NewInt64(tableName, "time_limit")
	_instNodeTask.ConnData = field.NewString(tableName, "conn_data")
	_instNodeTask.FormPerData = field.NewString(tableName, "form_per_data")
	_instNodeTask.Status = field.NewInt32(tableName, "status")
	_instNodeTask.CreateTime = field.NewTime(tableName, "create_time")
	_instNodeTask.UpdateTime = field.NewTime(tableName, "update_time")

	_instNodeTask.fillFieldMap()

	return _instNodeTask
}

type instNodeTask struct {
	instNodeTaskDo instNodeTaskDo

	ALL            field.Asterisk
	ID             field.Int64  // 唯一id
	InstTaskID     field.String // 实例任务id
	NodeTaskID     field.String // 节点任务id
	NodeID         field.String // 节点任务id
	ParentID       field.String // 父节点id
	NodeType       field.Int32  // 节点类型【1：正常节点；2：开始节点；3：结束节点；4：汇聚节点；5：条件节点；6：分支节点】
	NodeName       field.String // 节点名称
	ForwardMode    field.Int32  // 进行模式【1：并行 2：串行】
	CompleteConn   field.Int32  // 节点完成条件;通过的人数，0表示所有人通过，节点才算完成
	PermissionMode field.Int32  // 权限模式【1：协同 2：知会 3：审批；4：业务】
	AllowAdd       field.Int32  // 允许加签【1：不能加签；2：允许加签】
	ProcessMode    field.Int32  // 处理模式【1：人工； 2：自动；3：自动转人工】
	TimeLimit      field.Int64  // 处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限
	ConnData       field.String // 条件数据;前端生成，json格式
	FormPerData    field.String // 表单权限数据;节点表单权限配置，json格式
	Status         field.Int32  // 任务状态【0：未开始；1：处理中；2：完成；3：回退；4：终止；5：条件验证通过；6：条件验证不通过】
	CreateTime     field.Time   // 创建时间
	UpdateTime     field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (i instNodeTask) Table(newTableName string) *instNodeTask {
	i.instNodeTaskDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i instNodeTask) As(alias string) *instNodeTask {
	i.instNodeTaskDo.DO = *(i.instNodeTaskDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *instNodeTask) updateTableName(table string) *instNodeTask {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.InstTaskID = field.NewString(table, "inst_task_id")
	i.NodeTaskID = field.NewString(table, "node_task_id")
	i.NodeID = field.NewString(table, "node_id")
	i.ParentID = field.NewString(table, "parent_id")
	i.NodeType = field.NewInt32(table, "node_type")
	i.NodeName = field.NewString(table, "node_name")
	i.ForwardMode = field.NewInt32(table, "forward_mode")
	i.CompleteConn = field.NewInt32(table, "complete_conn")
	i.PermissionMode = field.NewInt32(table, "permission_mode")
	i.AllowAdd = field.NewInt32(table, "allow_add")
	i.ProcessMode = field.NewInt32(table, "process_mode")
	i.TimeLimit = field.NewInt64(table, "time_limit")
	i.ConnData = field.NewString(table, "conn_data")
	i.FormPerData = field.NewString(table, "form_per_data")
	i.Status = field.NewInt32(table, "status")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")

	i.fillFieldMap()

	return i
}

func (i *instNodeTask) WithContext(ctx context.Context) *instNodeTaskDo {
	return i.instNodeTaskDo.WithContext(ctx)
}

func (i instNodeTask) TableName() string { return i.instNodeTaskDo.TableName() }

func (i instNodeTask) Alias() string { return i.instNodeTaskDo.Alias() }

func (i *instNodeTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *instNodeTask) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 18)
	i.fieldMap["id"] = i.ID
	i.fieldMap["inst_task_id"] = i.InstTaskID
	i.fieldMap["node_task_id"] = i.NodeTaskID
	i.fieldMap["node_id"] = i.NodeID
	i.fieldMap["parent_id"] = i.ParentID
	i.fieldMap["node_type"] = i.NodeType
	i.fieldMap["node_name"] = i.NodeName
	i.fieldMap["forward_mode"] = i.ForwardMode
	i.fieldMap["complete_conn"] = i.CompleteConn
	i.fieldMap["permission_mode"] = i.PermissionMode
	i.fieldMap["allow_add"] = i.AllowAdd
	i.fieldMap["process_mode"] = i.ProcessMode
	i.fieldMap["time_limit"] = i.TimeLimit
	i.fieldMap["conn_data"] = i.ConnData
	i.fieldMap["form_per_data"] = i.FormPerData
	i.fieldMap["status"] = i.Status
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
}

func (i instNodeTask) clone(db *gorm.DB) instNodeTask {
	i.instNodeTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i instNodeTask) replaceDB(db *gorm.DB) instNodeTask {
	i.instNodeTaskDo.ReplaceDB(db)
	return i
}

type instNodeTaskDo struct{ gen.DO }

func (i instNodeTaskDo) Debug() *instNodeTaskDo {
	return i.withDO(i.DO.Debug())
}

func (i instNodeTaskDo) WithContext(ctx context.Context) *instNodeTaskDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i instNodeTaskDo) ReadDB() *instNodeTaskDo {
	return i.Clauses(dbresolver.Read)
}

func (i instNodeTaskDo) WriteDB() *instNodeTaskDo {
	return i.Clauses(dbresolver.Write)
}

func (i instNodeTaskDo) Session(config *gorm.Session) *instNodeTaskDo {
	return i.withDO(i.DO.Session(config))
}

func (i instNodeTaskDo) Clauses(conds ...clause.Expression) *instNodeTaskDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i instNodeTaskDo) Returning(value interface{}, columns ...string) *instNodeTaskDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i instNodeTaskDo) Not(conds ...gen.Condition) *instNodeTaskDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i instNodeTaskDo) Or(conds ...gen.Condition) *instNodeTaskDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i instNodeTaskDo) Select(conds ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i instNodeTaskDo) Where(conds ...gen.Condition) *instNodeTaskDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i instNodeTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *instNodeTaskDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i instNodeTaskDo) Order(conds ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i instNodeTaskDo) Distinct(cols ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i instNodeTaskDo) Omit(cols ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i instNodeTaskDo) Join(table schema.Tabler, on ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i instNodeTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i instNodeTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i instNodeTaskDo) Group(cols ...field.Expr) *instNodeTaskDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i instNodeTaskDo) Having(conds ...gen.Condition) *instNodeTaskDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i instNodeTaskDo) Limit(limit int) *instNodeTaskDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i instNodeTaskDo) Offset(offset int) *instNodeTaskDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i instNodeTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *instNodeTaskDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i instNodeTaskDo) Unscoped() *instNodeTaskDo {
	return i.withDO(i.DO.Unscoped())
}

func (i instNodeTaskDo) Create(values ...*model.InstNodeTask) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i instNodeTaskDo) CreateInBatches(values []*model.InstNodeTask, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i instNodeTaskDo) Save(values ...*model.InstNodeTask) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i instNodeTaskDo) First() (*model.InstNodeTask, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTask), nil
	}
}

func (i instNodeTaskDo) Take() (*model.InstNodeTask, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTask), nil
	}
}

func (i instNodeTaskDo) Last() (*model.InstNodeTask, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTask), nil
	}
}

func (i instNodeTaskDo) Find() ([]*model.InstNodeTask, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstNodeTask), err
}

func (i instNodeTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstNodeTask, err error) {
	buf := make([]*model.InstNodeTask, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i instNodeTaskDo) FindInBatches(result *[]*model.InstNodeTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i instNodeTaskDo) Attrs(attrs ...field.AssignExpr) *instNodeTaskDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i instNodeTaskDo) Assign(attrs ...field.AssignExpr) *instNodeTaskDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i instNodeTaskDo) Joins(fields ...field.RelationField) *instNodeTaskDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i instNodeTaskDo) Preload(fields ...field.RelationField) *instNodeTaskDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i instNodeTaskDo) FirstOrInit() (*model.InstNodeTask, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTask), nil
	}
}

func (i instNodeTaskDo) FirstOrCreate() (*model.InstNodeTask, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstNodeTask), nil
	}
}

func (i instNodeTaskDo) FindByPage(offset int, limit int) (result []*model.InstNodeTask, count int64, err error) {
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

func (i instNodeTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i instNodeTaskDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i instNodeTaskDo) Delete(models ...*model.InstNodeTask) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *instNodeTaskDo) withDO(do gen.Dao) *instNodeTaskDo {
	i.DO = *do.(*gen.DO)
	return i
}
