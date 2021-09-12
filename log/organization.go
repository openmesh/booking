package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
)

func OrganizationLoggingMiddleware(logger log.Logger) booking.OrganizationServiceMiddleware {
	return func(next booking.OrganizationService) booking.OrganizationService {
		return organizationLoggingMiddleware{logger, next}
	}
}

type organizationLoggingMiddleware struct {
	logger log.Logger
	booking.OrganizationService
}

func (mw organizationLoggingMiddleware) FindCurrentOrganization(ctx context.Context) (organization *booking.Organization, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_current_organization",
			"organization", organization,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	organization, err = mw.OrganizationService.FindCurrentOrganization(ctx)
	return
}

func (mw organizationLoggingMiddleware) FindOrganizationByPrivateKey(ctx context.Context, key string) (organization *booking.Organization, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_organization_by_private_key",
			"key", key,
			"organization", organization,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	organization, err = mw.OrganizationService.FindOrganizationByPrivateKey(ctx, key)
	return
}

func (mw organizationLoggingMiddleware) CreateOrganization(ctx context.Context, organization *booking.Organization) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_organization",
			"organization", organization,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.OrganizationService.CreateOrganization(ctx, organization)
	return
}

func (mw organizationLoggingMiddleware) UpdateOrganization(ctx context.Context, upd booking.OrganizationUpdate) (organization *booking.Organization, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_organization",
			"update", upd,
			"organization", organization,
			"err", err,
		)
	}(time.Now())

	organization, err = mw.OrganizationService.UpdateOrganization(ctx, upd)
	return
}
