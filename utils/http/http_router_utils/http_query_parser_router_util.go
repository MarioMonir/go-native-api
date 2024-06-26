package http_router_utils

import (
	"api/utils/regexp_utils"
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

// ------------------------------------------------------

func ParseQueryParamId(path string) (int, error) {
	matches := regexp.MustCompile(regexp_utils.GetQueryParamId).FindStringSubmatch(path)

	if len(matches) != 3 {
		return -1, errors.New("unable to parse id")
	}

	// Parse the id from the query params that matched
	id, err := strconv.Atoi(matches[2])
	if err != nil {
		return -1, err
	}

	return id, nil
}

// ------------------------------------------------------

func ParsePathValueId(r *http.Request) (int, error) {
	idString := r.PathValue("id")

	// Parse the id from the query params that matched
	id, err := strconv.Atoi(idString)
	if err != nil {
		return -1, err
	}

	return id, nil
}
