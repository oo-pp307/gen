package field

import "gorm.io/gorm/clause"

// Decimal Decimal type field
type Decimal Field

// Eq equal to
func (field Decimal) Eq(value Decimal) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to
func (field Decimal) Neq(value Decimal) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than
func (field Decimal) Gt(value Decimal) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to
func (field Decimal) Gte(value Decimal) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than
func (field Decimal) Lt(value Decimal) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to
func (field Decimal) Lte(value Decimal) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// In ...
func (field Decimal) In(values ...Decimal) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: field.toSlice(values...)}}
}

// NotIn ...
func (field Decimal) NotIn(values ...Decimal) Expr {
	return expr{e: clause.Not(field.In(values...).expression())}
}

// Between ...
func (field Decimal) Between(left Decimal, right Decimal) Expr {
	return field.between([]interface{}{left, right})
}

// NotBetween ...
func (field Decimal) NotBetween(left Decimal, right Decimal) Expr {
	return Not(field.Between(left, right))
}

// Like ...
func (field Decimal) Like(value Decimal) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// NotLike ...
func (field Decimal) NotLike(value Decimal) Expr {
	return expr{e: clause.Not(field.Like(value).expression())}
}

// Add ...
func (field Decimal) Add(value Decimal) Decimal {
	return Decimal{field.add(value)}
}

// Sub ...
func (field Decimal) Sub(value Decimal) Decimal {
	return Decimal{field.sub(value)}
}

// Mul ...
func (field Decimal) Mul(value Decimal) Decimal {
	return Decimal{field.mul(value)}
}

// Div ...
func (field Decimal) Div(value Decimal) Decimal {
	return Decimal{field.div(value)}
}

// FloorDiv ...
func (field Decimal) FloorDiv(value Decimal) Int {
	return Int{field.floorDiv(value)}
}

// Floor ...
func (field Decimal) Floor() Int {
	return Int{field.floor()}
}

// Value set value
func (field Decimal) Value(value Decimal) AssignExpr {
	return field.value(value)
}

// Zero set zero value
func (field Decimal) Zero() AssignExpr {
	return field.value(0)
}

// Sum ...
func (field Decimal) Sum() Decimal {
	return Decimal{field.sum()}
}

// IfNull ...
func (field Decimal) IfNull(value Decimal) Expr {
	return field.ifNull(value)
}

// Field ...
func (field Decimal) Field(values ...Decimal) Decimal {
	return Decimal{field.field(values)}
}

func (field Decimal) toSlice(values ...Decimal) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
