// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/sneaktricks/sport-matchmaking-match-service/model"
)

func newParticipation(db *gorm.DB, opts ...gen.DOOption) participation {
	_participation := participation{}

	_participation.participationDo.UseDB(db, opts...)
	_participation.participationDo.UseModel(&model.Participation{})

	tableName := _participation.participationDo.TableName()
	_participation.ALL = field.NewAsterisk(tableName)
	_participation.ID = field.NewUint(tableName, "id")
	_participation.CreatedAt = field.NewTime(tableName, "created_at")
	_participation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_participation.DeletedAt = field.NewField(tableName, "deleted_at")
	_participation.UserID = field.NewField(tableName, "user_id")
	_participation.MatchID = field.NewField(tableName, "match_id")

	_participation.fillFieldMap()

	return _participation
}

type participation struct {
	participationDo participationDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.Field
	MatchID   field.Field

	fieldMap map[string]field.Expr
}

func (p participation) Table(newTableName string) *participation {
	p.participationDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p participation) As(alias string) *participation {
	p.participationDo.DO = *(p.participationDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *participation) updateTableName(table string) *participation {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.UserID = field.NewField(table, "user_id")
	p.MatchID = field.NewField(table, "match_id")

	p.fillFieldMap()

	return p
}

func (p *participation) WithContext(ctx context.Context) *participationDo {
	return p.participationDo.WithContext(ctx)
}

func (p participation) TableName() string { return p.participationDo.TableName() }

func (p participation) Alias() string { return p.participationDo.Alias() }

func (p participation) Columns(cols ...field.Expr) gen.Columns {
	return p.participationDo.Columns(cols...)
}

func (p *participation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *participation) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["match_id"] = p.MatchID
}

func (p participation) clone(db *gorm.DB) participation {
	p.participationDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p participation) replaceDB(db *gorm.DB) participation {
	p.participationDo.ReplaceDB(db)
	return p
}

type participationDo struct{ gen.DO }

func (p participationDo) Debug() *participationDo {
	return p.withDO(p.DO.Debug())
}

func (p participationDo) WithContext(ctx context.Context) *participationDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p participationDo) ReadDB() *participationDo {
	return p.Clauses(dbresolver.Read)
}

func (p participationDo) WriteDB() *participationDo {
	return p.Clauses(dbresolver.Write)
}

func (p participationDo) Session(config *gorm.Session) *participationDo {
	return p.withDO(p.DO.Session(config))
}

func (p participationDo) Clauses(conds ...clause.Expression) *participationDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p participationDo) Returning(value interface{}, columns ...string) *participationDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p participationDo) Not(conds ...gen.Condition) *participationDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p participationDo) Or(conds ...gen.Condition) *participationDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p participationDo) Select(conds ...field.Expr) *participationDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p participationDo) Where(conds ...gen.Condition) *participationDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p participationDo) Order(conds ...field.Expr) *participationDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p participationDo) Distinct(cols ...field.Expr) *participationDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p participationDo) Omit(cols ...field.Expr) *participationDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p participationDo) Join(table schema.Tabler, on ...field.Expr) *participationDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p participationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *participationDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p participationDo) RightJoin(table schema.Tabler, on ...field.Expr) *participationDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p participationDo) Group(cols ...field.Expr) *participationDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p participationDo) Having(conds ...gen.Condition) *participationDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p participationDo) Limit(limit int) *participationDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p participationDo) Offset(offset int) *participationDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p participationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *participationDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p participationDo) Unscoped() *participationDo {
	return p.withDO(p.DO.Unscoped())
}

func (p participationDo) Create(values ...*model.Participation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p participationDo) CreateInBatches(values []*model.Participation, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p participationDo) Save(values ...*model.Participation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p participationDo) First() (*model.Participation, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Participation), nil
	}
}

func (p participationDo) Take() (*model.Participation, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Participation), nil
	}
}

func (p participationDo) Last() (*model.Participation, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Participation), nil
	}
}

func (p participationDo) Find() ([]*model.Participation, error) {
	result, err := p.DO.Find()
	return result.([]*model.Participation), err
}

func (p participationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Participation, err error) {
	buf := make([]*model.Participation, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p participationDo) FindInBatches(result *[]*model.Participation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p participationDo) Attrs(attrs ...field.AssignExpr) *participationDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p participationDo) Assign(attrs ...field.AssignExpr) *participationDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p participationDo) Joins(fields ...field.RelationField) *participationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p participationDo) Preload(fields ...field.RelationField) *participationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p participationDo) FirstOrInit() (*model.Participation, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Participation), nil
	}
}

func (p participationDo) FirstOrCreate() (*model.Participation, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Participation), nil
	}
}

func (p participationDo) FindByPage(offset int, limit int) (result []*model.Participation, count int64, err error) {
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

func (p participationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p participationDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p participationDo) Delete(models ...*model.Participation) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *participationDo) withDO(do gen.Dao) *participationDo {
	p.DO = *do.(*gen.DO)
	return p
}