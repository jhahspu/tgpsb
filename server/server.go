package server

import (
	"FS01/controller"
	"FS01/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	s := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowCredentials = true

	s.Use(gin.Recovery(), middleware.Logger(), cors.New(config))

	s.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := s.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/signup", controller.SignUp)
			public.POST("/login", controller.Login)
			public.GET("/posts/front", controller.PostsFront)
			public.GET("/posts/:stack", controller.PostsByStack)
			public.GET("/post/:id", controller.OnePost)
		}

		protected := api.Group("/protected").Use(middleware.CookieToken())
		{
			protected.GET("/user", controller.Profile)
			protected.POST("/user/post", controller.CreatePost)
			protected.GET("/user/post/:id", controller.GetPostForUser)
			protected.DELETE("/user/post/:id", controller.DeletePost)
			protected.POST("/logout", func(c *gin.Context) {
				c.SetCookie("jwt", "", -1, "/", "https://tpgs.herokuapp.com/", false, true)
			})
		}
	}

	return s
}
