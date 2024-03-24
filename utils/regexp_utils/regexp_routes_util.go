package regexp_utils

// Regular Experssions with Sprintf %s
const (
	GetListRegExp   = `^\/%s\/?$`
	GetOneRegExp    = `^\/%s\/(\d+)$`
	CreateOneRegExp = `^\/%s\/?$`
	UpdateOneRegExp = `^\/%s\/(\d+)$`
	DeleteOneRegExp = `^\/%s\/(\d+)$`

	GetQueryParamId = `^\/(\w*)\/(\d+)$`
)
