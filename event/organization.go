package event

import (
	"context"

	"github.com/openmesh/booking"
)

func OrganizationEventMiddleware(eventService booking.EventService) booking.OrganizationServiceMiddleware {
	return func(next booking.OrganizationService) booking.OrganizationService {
		return organizationEventMiddleware{eventService, next}
	}
}

type organizationEventMiddleware struct {
	booking.EventService
	booking.OrganizationService
}

// FindCurrentOrganization retrieves the organization of the currently authenticated user.
func (mw organizationEventMiddleware) FindCurrentOrganization(ctx context.Context) (*booking.Organization, error) {
	return mw.OrganizationService.FindCurrentOrganization(ctx)
}

// FindOrganizationByPrivateKey retrieves an organization by PrivateKey. Returns ENOTFOUND if
// organization does not exist.
func (mw organizationEventMiddleware) FindOrganizationByPrivateKey(ctx context.Context, key string) (*booking.Organization, error) {
	return mw.OrganizationService.FindOrganizationByPrivateKey(ctx, key)
}

// CreateOrganization creates a new organization.
func (mw organizationEventMiddleware) CreateOrganization(ctx context.Context, org *booking.Organization) (err error) {
	defer func() {
		if err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeOrganizationCreated,
			Payload: booking.OrganizationCreatedPayload{Organization: org},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	err = mw.OrganizationService.CreateOrganization(ctx, org)
	return
}

// UpdateOrganization updates the organization associated with the currently authenicated user.
func (mw organizationEventMiddleware) UpdateOrganization(ctx context.Context, upd booking.OrganizationUpdate) (org *booking.Organization, err error) {
	defer func() {
		if err != nil {
			return
		}
		ev := booking.Event{
			Type:    booking.EventTypeOrganizationUpdated,
			Payload: booking.OrganizationCreatedPayload{Organization: org},
		}
		userID := booking.UserIDFromContext(ctx)
		mw.EventService.PublishEvent(userID, ev)
	}()
	org, err = mw.OrganizationService.UpdateOrganization(ctx, upd)
	return
}
