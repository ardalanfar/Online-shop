package dto

import "Farashop/internal/entity"

//Member Contract Interactor Request And Response (Struct)

/*-------------------------Order Management----------------------------*/
/*--------------------Show Orders--------------------*/

type ShowOrdersRequest struct{}

type ShowOrdersResponse struct {
	Orders []entity.Order `json:"order"`
}

/*--------------------Edit Order---------------------*/

/*---------------------------------------------------*/
