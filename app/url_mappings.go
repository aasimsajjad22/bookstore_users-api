package app

import "github.com/aasimsajjad22/bookstore_users-api/controllers/ping"
import "github.com/aasimsajjad22/bookstore_users-api/controllers/users"

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.Get)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users", users.Create)
	router.POST("/users/login", users.Login)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

}
