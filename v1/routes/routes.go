package routes

func GetRoutes() []Router {
	var routes = append([]Router{}, GetTransactionRouters()[:]...)
	return routes
}

type Routes struct {
	Items []Router
}
