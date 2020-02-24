package handlers

import (
	"github.com/mateusprado/command-line-refactored/types"
)

func SliBuilder() types.SLI {

	var sli types.SLI
	sli.Id = "I09"
	sli.Name = "availabity indicator"

	return sli
}
