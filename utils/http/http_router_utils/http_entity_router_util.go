package http_router_utils

import (
	"api/utils/regexp_utils"
	"fmt"
	"net/http"
	"regexp"
)

type EntityRouter struct {
	req   *http.Request
	route string
}

func NewEntityRouter(req *http.Request, route string) *EntityRouter {
	return &EntityRouter{req, route}
}

func (er *EntityRouter) isMatchedRoute(pattern string) bool {
	return regexp.MustCompile(fmt.Sprintf(pattern, er.route)).MatchString(er.req.URL.Path)
}

func (er *EntityRouter) MatchGetListRoute() bool {
	if er.req.Method != http.MethodGet {
		return false
	}
	return er.isMatchedRoute(regexp_utils.GetListRegExp)
}

func (er *EntityRouter) MatchGetOneRoute() bool {
	if er.req.Method != http.MethodGet {
		return false
	}
	return er.isMatchedRoute(regexp_utils.GetOneRegExp)
}

func (er *EntityRouter) MatchCreateOneRoute() bool {
	if er.req.Method != http.MethodPost {
		return false
	}
	return er.isMatchedRoute(regexp_utils.CreateOneRegExp)
}

func (er *EntityRouter) MatchUpdateOneRoute() bool {
	if er.req.Method != http.MethodPut {
		return false
	}
	return er.isMatchedRoute(regexp_utils.UpdateOneRegExp)
}

func (er *EntityRouter) MatchDeleteOneRoute() bool {
	if er.req.Method != http.MethodDelete {
		return false
	}
	return er.isMatchedRoute(regexp_utils.UpdateOneRegExp)
}
