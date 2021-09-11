// Code generated by entc, DO NOT EDIT.

package resource

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmesh/booking/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Timezone applies equality check predicate on the "timezone" field. It's identical to TimezoneEQ.
func Timezone(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimezone), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// BookingPrice applies equality check predicate on the "bookingPrice" field. It's identical to BookingPriceEQ.
func BookingPrice(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBookingPrice), v))
	})
}

// OrganizationId applies equality check predicate on the "organizationId" field. It's identical to OrganizationIdEQ.
func OrganizationId(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrganizationId), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// TimezoneEQ applies the EQ predicate on the "timezone" field.
func TimezoneEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimezone), v))
	})
}

// TimezoneNEQ applies the NEQ predicate on the "timezone" field.
func TimezoneNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimezone), v))
	})
}

// TimezoneIn applies the In predicate on the "timezone" field.
func TimezoneIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimezone), v...))
	})
}

// TimezoneNotIn applies the NotIn predicate on the "timezone" field.
func TimezoneNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimezone), v...))
	})
}

// TimezoneGT applies the GT predicate on the "timezone" field.
func TimezoneGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimezone), v))
	})
}

// TimezoneGTE applies the GTE predicate on the "timezone" field.
func TimezoneGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimezone), v))
	})
}

// TimezoneLT applies the LT predicate on the "timezone" field.
func TimezoneLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimezone), v))
	})
}

// TimezoneLTE applies the LTE predicate on the "timezone" field.
func TimezoneLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimezone), v))
	})
}

// TimezoneContains applies the Contains predicate on the "timezone" field.
func TimezoneContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTimezone), v))
	})
}

// TimezoneHasPrefix applies the HasPrefix predicate on the "timezone" field.
func TimezoneHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTimezone), v))
	})
}

// TimezoneHasSuffix applies the HasSuffix predicate on the "timezone" field.
func TimezoneHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTimezone), v))
	})
}

// TimezoneEqualFold applies the EqualFold predicate on the "timezone" field.
func TimezoneEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTimezone), v))
	})
}

// TimezoneContainsFold applies the ContainsFold predicate on the "timezone" field.
func TimezoneContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTimezone), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// BookingPriceEQ applies the EQ predicate on the "bookingPrice" field.
func BookingPriceEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBookingPrice), v))
	})
}

// BookingPriceNEQ applies the NEQ predicate on the "bookingPrice" field.
func BookingPriceNEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBookingPrice), v))
	})
}

// BookingPriceIn applies the In predicate on the "bookingPrice" field.
func BookingPriceIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBookingPrice), v...))
	})
}

// BookingPriceNotIn applies the NotIn predicate on the "bookingPrice" field.
func BookingPriceNotIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBookingPrice), v...))
	})
}

// BookingPriceGT applies the GT predicate on the "bookingPrice" field.
func BookingPriceGT(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBookingPrice), v))
	})
}

// BookingPriceGTE applies the GTE predicate on the "bookingPrice" field.
func BookingPriceGTE(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBookingPrice), v))
	})
}

// BookingPriceLT applies the LT predicate on the "bookingPrice" field.
func BookingPriceLT(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBookingPrice), v))
	})
}

// BookingPriceLTE applies the LTE predicate on the "bookingPrice" field.
func BookingPriceLTE(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBookingPrice), v))
	})
}

// OrganizationIdEQ applies the EQ predicate on the "organizationId" field.
func OrganizationIdEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrganizationId), v))
	})
}

// OrganizationIdNEQ applies the NEQ predicate on the "organizationId" field.
func OrganizationIdNEQ(v int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrganizationId), v))
	})
}

// OrganizationIdIn applies the In predicate on the "organizationId" field.
func OrganizationIdIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrganizationId), v...))
	})
}

// OrganizationIdNotIn applies the NotIn predicate on the "organizationId" field.
func OrganizationIdNotIn(vs ...int) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrganizationId), v...))
	})
}

// HasSlots applies the HasEdge predicate on the "slots" edge.
func HasSlots() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SlotsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SlotsTable, SlotsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSlotsWith applies the HasEdge predicate on the "slots" edge with a given conditions (other predicates).
func HasSlotsWith(preds ...predicate.Slot) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SlotsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SlotsTable, SlotsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookings applies the HasEdge predicate on the "bookings" edge.
func HasBookings() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingsWith applies the HasEdge predicate on the "bookings" edge with a given conditions (other predicates).
func HasBookingsWith(preds ...predicate.Booking) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUnavailabilities applies the HasEdge predicate on the "unavailabilities" edge.
func HasUnavailabilities() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UnavailabilitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UnavailabilitiesTable, UnavailabilitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUnavailabilitiesWith applies the HasEdge predicate on the "unavailabilities" edge with a given conditions (other predicates).
func HasUnavailabilitiesWith(preds ...predicate.Unavailability) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UnavailabilitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UnavailabilitiesTable, UnavailabilitiesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
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
func Not(p predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		p(s.Not())
	})
}