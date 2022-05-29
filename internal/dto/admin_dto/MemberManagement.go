package admin_dto

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

type DeleteMemberResponse struct{}

/*------------------------------------------------*/
/*------------------------------------------------*/
