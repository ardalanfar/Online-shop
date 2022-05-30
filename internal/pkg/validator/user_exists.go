package validator

import (
	"Farashop/internal/contract"
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

func DoesUsernameExist(ctx context.Context, store contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		username := value.(string)

		ok, err := store.DoesUsernameExist(ctx, username)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if !ok {
			return fmt.Errorf("user: %s does not exist", username)
		}
		return nil
	}
}

func DoesIDExist(ctx context.Context, store contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		id := value.(uint)

		ok, err := store.DoesIDExist(ctx, id)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if !ok {
			return fmt.Errorf("id: %v does not exist", id)
		}
		return nil
	}
}
