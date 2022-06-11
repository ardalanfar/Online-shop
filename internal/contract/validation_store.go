package contract

import "context"

//validator Store package Contract (interface)

type ValidatorStore interface {
	DoesUsernameActiveStore(ctx context.Context, username string) (bool, error)
	DoesIDExistStore(ctx context.Context, id uint) (bool, error)
}
