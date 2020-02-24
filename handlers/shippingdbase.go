package handlers

import "github.com/mateusprado/shipping"

func RetrieveShippings() string {
	return shipping.RetrieveAllShippingInfo()
}
