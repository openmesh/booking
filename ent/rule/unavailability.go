package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/unavailability"
)

func FilterUnavailabilityOrganizationQueryRule() privacy.QueryRule {
	return privacy.UnavailabilityQueryRuleFunc(func(ctx context.Context, uq *ent.UnavailabilityQuery) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		uq.Where(unavailability.HasResourceWith(resource.OrganizationId(orgID)))
		return privacy.Skip
	})
}

func FilterUnavailabilityOrganizationMutationRule() privacy.MutationRule {
	return privacy.UnavailabilityMutationRuleFunc(func(ctx context.Context, um *ent.UnavailabilityMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		um.Where(unavailability.HasResourceWith(resource.OrganizationId(orgID)))
		return privacy.Skip
	})
}
