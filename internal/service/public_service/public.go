package public_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"Farashop/pkg/encrypt"
	"context"
	"math/rand"
)

type Interactor struct {
	store contract.PublicStore
}

func NewPublic(store contract.PublicStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) Register(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	//create hash password
	Password, errhash := encrypt.HashPassword(user.Password)
	if errhash != nil {
		return dto.RegisterUserResponse{Result: false}, errhash
	}
	user.Password = Password

	//create verification code
	min := 10000
	max := 99999
	randCode := rand.Intn(max-min) + min
	user.Verification_code = uint(randCode)

	//test
	// user.Verification_code = 55458
	// user.ID = 0

	//create user
	resCreate, errCrate := i.store.Register(ctx, user)
	if errCrate != nil || !resCreate {
		return dto.RegisterUserResponse{Result: false}, errCrate
	}
	//return
	return dto.RegisterUserResponse{Result: true}, nil
}

func (i Interactor) Login(ctx context.Context, req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Password: req.Password,
	}

	//get information user by username
	getInfo, errInfo := i.store.Login(ctx, user)
	if errInfo != nil {
		return dto.LoginUserResponse{}, errInfo
	}
	//check password with username
	checkpass := encrypt.CheckPasswordHash(user.Password, getInfo.Password)
	if checkpass != nil {
		return dto.LoginUserResponse{Result: false, User: getInfo}, checkpass
	}
	//return
	return dto.LoginUserResponse{Result: true, User: getInfo}, nil
}

func (i Interactor) MemberValidation(ctx context.Context, req dto.MemberValidationRequest) (dto.MemberValidationResponse, error) {
	user := entity.User{
		Username:          req.Username,
		Verification_code: req.Code,
	}

	//member validation
	resUpdate, errInfo := i.store.MemberValidation(ctx, user)
	if errInfo != nil || !resUpdate {
		return dto.MemberValidationResponse{Result: false}, errInfo
	}
	//return true
	return dto.MemberValidationResponse{Result: true}, nil
}
