package dto

//Member Contract Interactor Request And Response (Struct)

/*-------------------------Order Management----------------------------*/
/*--------------------Show Orders--------------------*/

type ShowOrdersRequest struct {
	ID uint `json:"id"`
}

type ShowOrdersResponse struct {
	Orders []Showorders `json:"orders"`
}

/*--------------------Edit Order---------------------*/
/*---------------------------------------------------*/
