package handlers

import "github.com/mateusprado/shipping"

func ListAllFromHttpService() string {

	return shipping.GetShippingInfo()

}
