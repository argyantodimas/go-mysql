package controllers

import "github.com/argyantodimas/go-mysql/api/middlewares"

func (app *App) initializeRoutes() {

	// Home Route
	app.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(app.Home)).Methods("GET")

	// Login Route
	app.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(app.Login)).Methods("POST")

	//Users routes
	app.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(app.CreateUser)).Methods("POST")
	app.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(app.GetUsers)).Methods("GET")
	app.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(app.GetUser)).Methods("GET")
	app.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(app.UpdateUser))).Methods("PUT")
	app.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(app.DeleteUser)).Methods("DELETE")

	//Posts routes
	app.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(app.CreatePost)).Methods("POST")
	app.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(app.GetPosts)).Methods("GET")
	app.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(app.GetPost)).Methods("GET")
	app.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(app.UpdatePost))).Methods("PUT")
	app.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(app.DeletePost)).Methods("DELETE")
}
