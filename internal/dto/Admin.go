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
type ShowInfoMemberRequest struct {
	ID uint `json:"id"`
}

type ShowInfoMemberResponse struct {
	User   entity.User
}

/*------------------------------------------------*/
/*--------------------------------------------------------------------------*/

/*--------------------------------------------------------------------------*/
