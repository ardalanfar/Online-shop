package validator

import (
	"Farashop/internal/contract"
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

func doesUserExist(ctx context.Context, store contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		username := value.(string)

		ok, err := store.DoesUserExist(ctx, username)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if !ok {
			return fmt.Errorf("user: %d does not exist", username)
		}

		return nil
	}
}
