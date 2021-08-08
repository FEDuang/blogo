package api

import (
	blog "blogo/api/blog"
	"blogo/api/system"
)

type ApiGroup struct {
	BlogApiGroup		blog.ApiGroup
	SystemApiGroup		system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
