// Code generated by ent, DO NOT EDIT.

package page

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tomato3713/nullwiki/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldID, id))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldBody, v))
}

// TextFormat applies equality check predicate on the "text_format" field. It's identical to TextFormatEQ.
func TextFormat(v uint) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldTextFormat, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldUpdatedAt, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldBody, v))
}

// TextFormatEQ applies the EQ predicate on the "text_format" field.
func TextFormatEQ(v uint) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldTextFormat, v))
}

// TextFormatNEQ applies the NEQ predicate on the "text_format" field.
func TextFormatNEQ(v uint) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldTextFormat, v))
}

// TextFormatIn applies the In predicate on the "text_format" field.
func TextFormatIn(vs ...uint) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldTextFormat, vs...))
}

// TextFormatNotIn applies the NotIn predicate on the "text_format" field.
func TextFormatNotIn(vs ...uint) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldTextFormat, vs...))
}

// TextFormatGT applies the GT predicate on the "text_format" field.
func TextFormatGT(v uint) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldTextFormat, v))
}

// TextFormatGTE applies the GTE predicate on the "text_format" field.
func TextFormatGTE(v uint) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldTextFormat, v))
}

// TextFormatLT applies the LT predicate on the "text_format" field.
func TextFormatLT(v uint) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldTextFormat, v))
}

// TextFormatLTE applies the LTE predicate on the "text_format" field.
func TextFormatLTE(v uint) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldTextFormat, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		p(s.Not())
	})
}
