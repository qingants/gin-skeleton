package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qingants/gin-skeleton/api/v1"
)

func movie(app *gin.Engine) {
	m := app.Group("/movie/v1")
	{
		m.GET("/movies", v1.GetMovies)
		m.POST("/love/add", v1.AddLoveMovie)
		m.POST("/love/del", v1.DelLoveMovie)
		m.GET("/love/movies", v1.GetLoveMovie)
	}
}

