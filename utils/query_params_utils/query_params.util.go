package query_params_utils

import (
	"fmt"
	"regexp"
)

func RegexpToGetEntityById(entityRouteName string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^\/%s\/(\d+)$`, entityRouteName))
}
