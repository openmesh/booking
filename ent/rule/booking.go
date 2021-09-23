package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	entbooking "github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
)

func FilterBookingOrganizationQueryRule() privacy.QueryRule {
	return privacy.BookingQueryRuleFunc(func(ctx context.Context, bq *ent.BookingQuery) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		bq.Where(
			entbooking.HasResourceWith(
				resource.OrganizationId(orgID),
			),
		)
		return privacy.Skip
	})
}

func FilterBookingOrganizationMutationRule() privacy.MutationRule {
	return privacy.BookingMutationRuleFunc(func(ctx context.Context, bm *ent.BookingMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		bm.Where(
			entbooking.HasResourceWith(
				resource.OrganizationId(orgID),
			),
		)
		return privacy.Skip
	})
}
