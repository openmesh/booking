package booking

import (
	"context"
	"fmt"
)

type Cache interface {
	Get(ctx context.Context, key string, val interface{}) error
	Refresh(ctx context.Context, key string) error
	Remove(ctx context.Context, key string) error
	RemoveMany(ctx context.Context, match string) error
	Set(ctx context.Context, key string, val interface{}) error
}

// CacheKeyFormat is the standard format for cache keys. It should be in the
// format "openmesh:booking:{method}:{organizationID}:{params}"
const CacheKeyFormat = "openmesh:booking:%s:%d:%+v"

func CacheKey(ctx context.Context, method string, req interface{}) string {
	orgID := OrganizationIDFromContext(ctx)
	return fmt.Sprintf(
		CacheKeyFormat,
		method,
		orgID,
		req,
	)
}
