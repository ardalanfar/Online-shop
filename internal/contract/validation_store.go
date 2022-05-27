package contract

import "context"

type ValidatorStore interface {
	DoesUserExist(ctx context.Context, username string) (bool, error)
}
