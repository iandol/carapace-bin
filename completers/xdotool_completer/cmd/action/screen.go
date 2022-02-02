package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionScreens() carapace.Action {
	return carapace.ActionExecCommand("xrandr")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<screen>[^ ]+) connected .*`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
