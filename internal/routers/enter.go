package routers

import (
	"goecommerce/internal/routers/manager"
	"goecommerce/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
