package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/token"
)

func FilterTokenUserMutationRule() privacy.MutationRule {
	return privacy.TokenMutationRuleFunc(func(ctx context.Context, tm *ent.TokenMutation) error {
		userID := booking.UserIDFromContext(ctx)
		if userID == 0 {
			return privacy.Denyf("missing user from context")
		}
		tm.Where(token.UserId(userID))
		return privacy.Skip
	})
}

func FilterTokenOrganizationMutationRule() privacy.MutationRule {
	return privacy.TokenMutationRuleFunc(func(ctx context.Context, tm *ent.TokenMutation) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization from context")
		}
		tm.Where(token.OrganizationId(orgID))
		return privacy.Skip
	})
}
