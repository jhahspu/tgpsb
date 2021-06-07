package server

import (
	"FS01/controller"
	"FS01/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	s := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200", "*"}
	config.AllowMethods = []string{"Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE"}
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}

	s.Use(gin.Recovery(), middleware.Logger(), cors.New(config))

	s.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	s.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	s.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := s.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/signup", controller.SignUp)
			public.POST("/sign-in", controller.Login)
			public.GET("/posts/front", controller.PostsFront)
			public.GET("/posts/:stack", controller.PostsByStack)
			public.GET("/post/:id", controller.OnePost)
		}

		protected := api.Group("/protected").Use(middleware.CookieToken())
		{
			protected.GET("/user", controller.Profile)
			protected.GET("/user/posts", controller.GetPostsForUser)
			protected.POST("/user/post", controller.CreatePost)
			protected.GET("/user/post/:id", controller.GetPostForUser)
			protected.DELETE("/user/post/:id", controller.DeletePost)
			protected.GET("/sign-out", controller.SignOut)
		}
	}

	return s
}
