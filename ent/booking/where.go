// Code generated by entc, DO NOT EDIT.

package booking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmesh/booking/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// ResourceId applies equality check predicate on the "resourceId" field. It's identical to ResourceIdEQ.
func ResourceId(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceId), v))
	})
}

// OrganizationId applies equality check predicate on the "organizationId" field. It's identical to OrganizationIdEQ.
func OrganizationId(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrganizationId), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...time.Time) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// ResourceIdEQ applies the EQ predicate on the "resourceId" field.
func ResourceIdEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceId), v))
	})
}

// ResourceIdNEQ applies the NEQ predicate on the "resourceId" field.
func ResourceIdNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResourceId), v))
	})
}

// ResourceIdIn applies the In predicate on the "resourceId" field.
func ResourceIdIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResourceId), v...))
	})
}

// ResourceIdNotIn applies the NotIn predicate on the "resourceId" field.
func ResourceIdNotIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResourceId), v...))
	})
}

// OrganizationIdEQ applies the EQ predicate on the "organizationId" field.
func OrganizationIdEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrganizationId), v))
	})
}

// OrganizationIdNEQ applies the NEQ predicate on the "organizationId" field.
func OrganizationIdNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrganizationId), v))
	})
}

// OrganizationIdIn applies the In predicate on the "organizationId" field.
func OrganizationIdIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
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
func OrganizationIdNotIn(vs ...int) predicate.Booking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrganizationId), v...))
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.BookingMetadatum) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResource applies the HasEdge predicate on the "resource" edge.
func HasResource() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResourceTable, ResourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourceWith applies the HasEdge predicate on the "resource" edge with a given conditions (other predicates).
func HasResourceWith(preds ...predicate.Resource) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResourceTable, ResourceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
