package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	entbooking "github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/bookingmetadatum"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
)

func FilterBookingMetadatumOrganizationQueryRule() privacy.QueryRule {
	return privacy.BookingMetadatumQueryRuleFunc(func(ctx context.Context, bmq *ent.BookingMetadatumQuery) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		bmq.Where(
			bookingmetadatum.HasBookingWith(
				entbooking.HasResourceWith(
					resource.OrganizationId(orgID),
				),
			),
		)
		return privacy.Skip
	})
}

func FilterBookingMetadatumOrganizationMutationRule() privacy.MutationRule {
	return privacy.BookingMetadatumMutationRuleFunc(func(ctx context.Context, bmm *ent.BookingMetadatumMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		bmm.Where(
			bookingmetadatum.HasBookingWith(
				entbooking.HasResourceWith(
					resource.OrganizationId(orgID),
				),
			),
		)
		return privacy.Skip
	})
}
