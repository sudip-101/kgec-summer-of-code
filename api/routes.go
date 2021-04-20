package api

func (server *Server) initRoutes() {
	g := GitHubAPI{}
	g.Init()
	server.Router.GET("/", g.RenderHome)
	server.Router.GET("/refresh", g.Refresh)
	server.Router.GET("/dashboard", g.RenderDashboard)
}
