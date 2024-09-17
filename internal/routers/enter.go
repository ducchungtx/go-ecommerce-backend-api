package routers

import "goecommerce/internal/routers/user"

type RouterGroup struct {
	User user.UserRouterGroup
}
