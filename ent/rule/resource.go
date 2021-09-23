package rule

import (
	"context"

	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/organization"
	"github.com/openmesh/booking/ent/predicate"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/resource"
)

func FilterOrganizationRule() privacy.QueryMutationRule {
	// OrganizationsFilter is an interface to wrap WhereHasOrganizationWith()
	// predicate that is used by protected schemas.
	type OrganizationsFilter interface {
		WhereHasOrganizationWith(...predicate.Organization)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		orgID := booking.OrganizationIDFromContext(ctx)
		if orgID == 0 {
			return privacy.Denyf("missing organization information")
		}
		of, ok := f.(OrganizationsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		of.WhereHasOrganizationWith(organization.ID(orgID))
		// Skip to the next privacy rule.
		return privacy.Skip
	})
}

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

// DenyMismatchedTenants is a rule that runs only on create operations and
// returns a deny decision if the operation tries to create entities that with
// organization IDs that are not valid for the current authenticated user.

// func DenyMismatchedOrganizations() privacy.MutationRule {
// 	return privacy.ResourceQueryRuleFunc()
// }
