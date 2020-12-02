package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kiwisheets/auth"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if bypass {
		return next(ctx)
	}

	if auth.For(ctx).UserID == 0 {
		return nil, fmt.Errorf("not logged in")
	}
	return next(ctx)
}

func IsSecureAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if bypass {
		return next(ctx)
	}

	if auth.For(ctx).UserID == 0 {
		return nil, fmt.Errorf("not logged in")
	}
	if !auth.For(ctx).Secure {
		return nil, fmt.Errorf("not logged in with time sensitive token")
	}
	return next(ctx)
}

func HasPerm(ctx context.Context, obj interface{}, next graphql.Resolver, perm string) (res interface{}, err error) {
	if bypass {
		return next(ctx)
	}

	if auth.For(ctx).UserID == 0 {
		return nil, fmt.Errorf("not logged in")
	}

	for _, p := range auth.For(ctx).Scopes {
		if p.CheckPermissionString(perm) {
			return next(ctx)
		}
	}

	return nil, fmt.Errorf("not authorised")
}

func HasPerms(ctx context.Context, obj interface{}, next graphql.Resolver, requestedPerms []string) (res interface{}, err error) {
	if bypass {
		return next(ctx)
	}

	if auth.For(ctx).UserID == 0 {
		return nil, fmt.Errorf("not logged in")
	}

	permsPassed := make([]bool, len(requestedPerms))

	for _, userPerm := range auth.For(ctx).Scopes {
		for i, requestedPerm := range requestedPerms {
			if permsPassed[i] {
				continue
			}
			if userPerm.CheckPermissionString(requestedPerm) {
				permsPassed[i] = true
			}
		}
	}
	for _, p := range permsPassed {
		if !p {
			return nil, fmt.Errorf("not authorised")
		}
	}

	return next(ctx)
}
