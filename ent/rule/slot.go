package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/slot"
)

func FilterSlotOrganizationQueryRule() privacy.QueryRule {
	return privacy.SlotQueryRuleFunc(func(ctx context.Context, sq *ent.SlotQuery) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		sq.Where(slot.HasResourceWith(resource.OrganizationId(orgID)))
		return privacy.Skip
	})
}

func FilterSlotOrganizationMutationRule() privacy.MutationRule {
	return privacy.SlotMutationRuleFunc(func(ctx context.Context, sm *ent.SlotMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		sm.Where(slot.HasResourceWith(resource.OrganizationId(orgID)))
		return privacy.Skip
	})
}
