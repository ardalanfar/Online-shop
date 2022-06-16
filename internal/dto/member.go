package dto

//Member Contract Interactor Request And Response (Struct)

/*-------------------------Order Management----------------------------*/
/*--------------------Show Orders--------------------*/

type ShowOrdersRequest struct{}

type ShowOrdersResponse struct {
	Orders []Showorders `json:"orders"`
}

/*--------------------Edit Order---------------------*/

/*---------------------------------------------------*/
