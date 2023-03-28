// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                     db,
		FlowDefInfo:            newFlowDefInfo(db, opts...),
		FlowDefNodeHandler:     newFlowDefNodeHandler(db, opts...),
		FlowDefNodeInfo:        newFlowDefNodeInfo(db, opts...),
		FormDefDetail:          newFormDefDetail(db, opts...),
		FormDefElement:         newFormDefElement(db, opts...),
		InstHandlerTask:        newInstHandlerTask(db, opts...),
		InstHandlerTaskOpinion: newInstHandlerTaskOpinion(db, opts...),
		InstNodeTask:           newInstNodeTask(db, opts...),
		InstTaskDetail:         newInstTaskDetail(db, opts...),
		InstTaskOpLog:          newInstTaskOpLog(db, opts...),
		InstTaskParam:          newInstTaskParam(db, opts...),
		InstTaskParamAttr:      newInstTaskParamAttr(db, opts...),
		ModelAuth:              newModelAuth(db, opts...),
		ModelDetail:            newModelDetail(db, opts...),
		ModelGroup:             newModelGroup(db, opts...),
		ModelVersion:           newModelVersion(db, opts...),
		OrganizationInfo:       newOrganizationInfo(db, opts...),
		RoleInfo:               newRoleInfo(db, opts...),
		UserInfo:               newUserInfo(db, opts...),
		UserRoleLink:           newUserRoleLink(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	FlowDefInfo            flowDefInfo
	FlowDefNodeHandler     flowDefNodeHandler
	FlowDefNodeInfo        flowDefNodeInfo
	FormDefDetail          formDefDetail
	FormDefElement         formDefElement
	InstHandlerTask        instHandlerTask
	InstHandlerTaskOpinion instHandlerTaskOpinion
	InstNodeTask           instNodeTask
	InstTaskDetail         instTaskDetail
	InstTaskOpLog          instTaskOpLog
	InstTaskParam          instTaskParam
	InstTaskParamAttr      instTaskParamAttr
	ModelAuth              modelAuth
	ModelDetail            modelDetail
	ModelGroup             modelGroup
	ModelVersion           modelVersion
	OrganizationInfo       organizationInfo
	RoleInfo               roleInfo
	UserInfo               userInfo
	UserRoleLink           userRoleLink
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		FlowDefInfo:            q.FlowDefInfo.clone(db),
		FlowDefNodeHandler:     q.FlowDefNodeHandler.clone(db),
		FlowDefNodeInfo:        q.FlowDefNodeInfo.clone(db),
		FormDefDetail:          q.FormDefDetail.clone(db),
		FormDefElement:         q.FormDefElement.clone(db),
		InstHandlerTask:        q.InstHandlerTask.clone(db),
		InstHandlerTaskOpinion: q.InstHandlerTaskOpinion.clone(db),
		InstNodeTask:           q.InstNodeTask.clone(db),
		InstTaskDetail:         q.InstTaskDetail.clone(db),
		InstTaskOpLog:          q.InstTaskOpLog.clone(db),
		InstTaskParam:          q.InstTaskParam.clone(db),
		InstTaskParamAttr:      q.InstTaskParamAttr.clone(db),
		ModelAuth:              q.ModelAuth.clone(db),
		ModelDetail:            q.ModelDetail.clone(db),
		ModelGroup:             q.ModelGroup.clone(db),
		ModelVersion:           q.ModelVersion.clone(db),
		OrganizationInfo:       q.OrganizationInfo.clone(db),
		RoleInfo:               q.RoleInfo.clone(db),
		UserInfo:               q.UserInfo.clone(db),
		UserRoleLink:           q.UserRoleLink.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		FlowDefInfo:            q.FlowDefInfo.replaceDB(db),
		FlowDefNodeHandler:     q.FlowDefNodeHandler.replaceDB(db),
		FlowDefNodeInfo:        q.FlowDefNodeInfo.replaceDB(db),
		FormDefDetail:          q.FormDefDetail.replaceDB(db),
		FormDefElement:         q.FormDefElement.replaceDB(db),
		InstHandlerTask:        q.InstHandlerTask.replaceDB(db),
		InstHandlerTaskOpinion: q.InstHandlerTaskOpinion.replaceDB(db),
		InstNodeTask:           q.InstNodeTask.replaceDB(db),
		InstTaskDetail:         q.InstTaskDetail.replaceDB(db),
		InstTaskOpLog:          q.InstTaskOpLog.replaceDB(db),
		InstTaskParam:          q.InstTaskParam.replaceDB(db),
		InstTaskParamAttr:      q.InstTaskParamAttr.replaceDB(db),
		ModelAuth:              q.ModelAuth.replaceDB(db),
		ModelDetail:            q.ModelDetail.replaceDB(db),
		ModelGroup:             q.ModelGroup.replaceDB(db),
		ModelVersion:           q.ModelVersion.replaceDB(db),
		OrganizationInfo:       q.OrganizationInfo.replaceDB(db),
		RoleInfo:               q.RoleInfo.replaceDB(db),
		UserInfo:               q.UserInfo.replaceDB(db),
		UserRoleLink:           q.UserRoleLink.replaceDB(db),
	}
}

type queryCtx struct {
	FlowDefInfo            *flowDefInfoDo
	FlowDefNodeHandler     *flowDefNodeHandlerDo
	FlowDefNodeInfo        *flowDefNodeInfoDo
	FormDefDetail          *formDefDetailDo
	FormDefElement         *formDefElementDo
	InstHandlerTask        *instHandlerTaskDo
	InstHandlerTaskOpinion *instHandlerTaskOpinionDo
	InstNodeTask           *instNodeTaskDo
	InstTaskDetail         *instTaskDetailDo
	InstTaskOpLog          *instTaskOpLogDo
	InstTaskParam          *instTaskParamDo
	InstTaskParamAttr      *instTaskParamAttrDo
	ModelAuth              *modelAuthDo
	ModelDetail            *modelDetailDo
	ModelGroup             *modelGroupDo
	ModelVersion           *modelVersionDo
	OrganizationInfo       *organizationInfoDo
	RoleInfo               *roleInfoDo
	UserInfo               *userInfoDo
	UserRoleLink           *userRoleLinkDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		FlowDefInfo:            q.FlowDefInfo.WithContext(ctx),
		FlowDefNodeHandler:     q.FlowDefNodeHandler.WithContext(ctx),
		FlowDefNodeInfo:        q.FlowDefNodeInfo.WithContext(ctx),
		FormDefDetail:          q.FormDefDetail.WithContext(ctx),
		FormDefElement:         q.FormDefElement.WithContext(ctx),
		InstHandlerTask:        q.InstHandlerTask.WithContext(ctx),
		InstHandlerTaskOpinion: q.InstHandlerTaskOpinion.WithContext(ctx),
		InstNodeTask:           q.InstNodeTask.WithContext(ctx),
		InstTaskDetail:         q.InstTaskDetail.WithContext(ctx),
		InstTaskOpLog:          q.InstTaskOpLog.WithContext(ctx),
		InstTaskParam:          q.InstTaskParam.WithContext(ctx),
		InstTaskParamAttr:      q.InstTaskParamAttr.WithContext(ctx),
		ModelAuth:              q.ModelAuth.WithContext(ctx),
		ModelDetail:            q.ModelDetail.WithContext(ctx),
		ModelGroup:             q.ModelGroup.WithContext(ctx),
		ModelVersion:           q.ModelVersion.WithContext(ctx),
		OrganizationInfo:       q.OrganizationInfo.WithContext(ctx),
		RoleInfo:               q.RoleInfo.WithContext(ctx),
		UserInfo:               q.UserInfo.WithContext(ctx),
		UserRoleLink:           q.UserRoleLink.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
