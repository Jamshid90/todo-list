package postgres

import (
	sq "github.com/Masterminds/squirrel"
)

// Squirrel provides wrapper around squirrel package
type Squirrel struct {
	Builder sq.StatementBuilderType
}

func NewSquirrel() *Squirrel {
	return &Squirrel{sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (s *Squirrel) Equal(key string, value interface{}) sq.Eq {
	return sq.Eq{key: value}
}

func (s *Squirrel) EqualStr(key string) EqualStr {
	return EqualStr(key)
}

func (s *Squirrel) ILike(key string, value interface{}) sq.ILike {
	return sq.ILike{key: value}
}

func (s *Squirrel) NotEqual(key string, value interface{}) sq.NotEq {
	return sq.NotEq{key: value}
}

func (s *Squirrel) Or(cond ...sq.Sqlizer) sq.Or {
	sl := make([]sq.Sqlizer, 0, len(cond))
	for _, val := range cond {
		sl = append(sl, val)
	}
	return sl
}

func (s *Squirrel) And(cond ...sq.Sqlizer) sq.And {
	sl := make([]sq.Sqlizer, 0, len(cond))
	for _, val := range cond {
		sl = append(sl, val)
	}
	return sl
}

func (s *Squirrel) Alias(expr sq.Sqlizer, alias string) sq.Sqlizer {
	return sq.Alias(expr, alias)
}

func (s *Squirrel) EqualMany(clauses map[string]interface{}) sq.Eq {
	eqMany := make(sq.Eq, len(clauses))
	for key, value := range clauses {
		eqMany[key] = value
	}
	return eqMany
}

func (s *Squirrel) Gt(key string, value interface{}) sq.Gt {
	return sq.Gt{key: value}
}

func (s *Squirrel) GtOrEq(key string, value interface{}) sq.GtOrEq {
	return sq.GtOrEq{key: value}
}

func (s *Squirrel) Lt(key string, value interface{}) sq.Lt {
	return sq.Lt{key: value}
}

func (s *Squirrel) LtOrEq(key string, value interface{}) sq.LtOrEq {
	return sq.LtOrEq{key: value}
}

func (s *Squirrel) Expr(sql string, args ...interface{}) sq.Sqlizer {
	return sq.Expr(sql, args)
}

type EqualStr string

func (e EqualStr) ToSql() (sql string, args []interface{}, err error) {
	sql = string(e)
	return
}
