package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
)

func FilterResourceOrganizationQueryRule() privacy.QueryRule {
	return privacy.ResourceQueryRuleFunc(func(ctx context.Context, rq *ent.ResourceQuery) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		rq.Where(resource.OrganizationId(orgID))
		return privacy.Skip
	})
}

func FilterResourceOrganizationMutationRule() privacy.MutationRule {
	return privacy.ResourceMutationRuleFunc(func(ctx context.Context, rm *ent.ResourceMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		rm.Where(resource.OrganizationId(orgID))
		return privacy.Skip
	})
}
