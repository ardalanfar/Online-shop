package contract

import "context"

//Validator Store package Contract (interface)

type ValidatorStore interface {
	DoesUsernameExist(ctx context.Context, username string) (bool, error)
	DoesIDExist(ctx context.Context, id uint) (bool, error)
}
