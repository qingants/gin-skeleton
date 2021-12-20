package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qingants/gin-skeleton/model"
	"github.com/qingants/gin-skeleton/pkg/errs"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type movie struct {
	ID uint
	Name string
}

//movies := []movie{}
//uid := 100

//func GetFakeUid() uint {
//	return 10000
//}

func GetFakeMovies() map[uint]movie {
	m := map[uint]movie{}

	m[1] = movie{
		ID:   1,
		Name: "movie1",
	}

	m[2] = movie{
		ID:   2,
		Name: "movie2",
	}

	m[3] = movie{
		ID:   3,
		Name: "movie3",
	}
	return m
}


func GetMovies(c *gin.Context)  {
	c.JSON(http.StatusOK, errs.GetResponse(errs.Success, GetFakeMovies()))
}

func GetUid(c *gin.Context) (error, uint) {
	uid := c.GetHeader("uid")
	u, err := strconv.Atoi(uid)
	if err != nil {
		//zap.L().Error(err.Error())
		//c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return err, 0
	}
	return nil, uint(u)
}

func AddLoveMovie(c *gin.Context)  {
	l := zap.L()

	mid := c.Query("mid")

	err, u := GetUid(c)
	if err != nil {
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	// mid check
	m, err := strconv.Atoi(mid)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	if _, ok := GetFakeMovies()[uint(m)]; !ok {
		l.Warn("not exists mid", zap.String("mid", mid))
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	result := model.Get().Create(&model.LoveMovie{Uid: uint(u), Mid: uint(m)})
	if result.Error != nil {
		l.Warn("db error", zap.Error(result.Error), zap.Int("mid", m), zap.Uint("uid", u))
		c.JSON(http.StatusOK, errs.GetResponse(errs.Maintenance, nil))
		return
	}

	c.JSON(http.StatusOK, errs.GetResponse(errs.Success, nil))
}

func DelLoveMovie(c *gin.Context)  {
	mid := c.Query("mid")
	l := zap.L()


	err, u := GetUid(c)
	if err != nil {
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	// mid check
	m, err := strconv.Atoi(mid)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	if _, ok := GetFakeMovies()[uint(m)]; !ok {
		l.Warn("not exists mid", zap.String("mid", mid))
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	result := model.Get().Where("mid = ? and uid = ?", uint(m), u).Delete(&model.LoveMovie{})
	if result.Error != nil {
		l.Warn("db error", zap.Error(result.Error), zap.Int("mid", m), zap.Uint("uid", u))
		c.JSON(http.StatusOK, errs.GetResponse(errs.Maintenance, nil))
		return
	}
	c.JSON(http.StatusOK, errs.GetResponse(errs.Success, nil))
}

func GetLoveMovie(c *gin.Context)  {

	err, u := GetUid(c)
	if err != nil {
		c.JSON(http.StatusOK, errs.GetResponse(errs.InvalidParams, nil))
		return
	}

	var movies []movie
	var loves []model.LoveMovie
	result := model.Get().Where("uid = ?", u).Find(&loves)
	if result.Error != nil {
		zap.L().Warn(result.Error.Error())
		c.JSON(http.StatusOK, errs.GetResponse(errs.Maintenance, nil))
		return
	}

	for _, l := range loves {
		if m, ok := GetFakeMovies()[l.Mid]; ok {
			movies = append(movies, m)
		}
	}

	c.JSON(http.StatusOK, errs.GetResponse(errs.Success, movies))
}