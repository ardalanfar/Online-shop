package dto

//Member Management Contract Interactor Request And Response (Struct)

import "Farashop/internal/entity"

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
