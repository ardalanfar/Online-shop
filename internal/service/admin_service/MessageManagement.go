package admin_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"context"
)

type Interactormsq struct {
	store contract.AdminStore
}

func NewMsg(store contract.AdminStore) Interactormsq {
	return Interactormsq{store: store}
}

func (i Interactormsq) SendMsg(ctx context.Context, req dto.SendMsgRequest) (dto.SendMsgResponse, error) {
	user := entity.User{
		ID: req.ID,
	}

	//get information user by username
	getInfo, errInfo := i.store.SendMsg(ctx, user)
	if errInfo != nil {
		return dto.SendMsgResponse{Result: false}, errInfo
	}
	//return true
	return dto.SendMsgResponse{Result: true, User: getInfo}, nil

}
