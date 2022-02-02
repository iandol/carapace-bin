package action

import "github.com/rsteube/carapace"

func ActionMouseButtons() carapace.Action {
	return carapace.ActionValuesDescribed(
		"1", "left mouse",
		"2", "middle mouse",
		"3", "right mouse",
		"4", "wheel up",
		"5", "wheel down",
	)
}
