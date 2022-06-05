package dto

//Admin Contract Interactor Request And Response (Struct)

import "Farashop/internal/entity"

/*-------------------------Member Management----------------------------*/
/*----------------ShowMembers--------------------*/

type ShowMembersRequest struct{}

type ShowMembersResponse struct {
	Users []entity.User `json:"users"`
}

/*---------------DeleteMember---------------------*/

type DeleteMemberRequest struct {
	ID uint `json:"id"`
}

type DeleteMemberResponse struct {
	Result bool `json:"result"`
}

/*------------------------------------------------*/
/*---------------------------Message Management-----------------------------*/
/*----------------SendMassage--------------------*/
type SendMsgRequest struct {
	ID uint `json:"id"`
}

type SendMsgResponse struct {
	User   entity.User
	Result bool `json:"result"`
}

/*------------------------------------------------*/
/*--------------------------------------------------------------------------*/

/*--------------------------------------------------------------------------*/
