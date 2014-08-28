package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Pagination struct {
	SinceId int64
	MaxId   int64
	Count   int
}

func PageMiddleware(errCode int, err interface{}) func(
	martini.Context, *http.Request, render.Render) {
	return func(c martini.Context, req *http.Request, r render.Render) {
		qs := req.URL.Query()
		var (
			page Pagination
			err  error
		)

		parseInt64 := func(str string) (i int64, err error) {
			if str == "" {
				return 0, nil
			}
			i, err = strconv.ParseInt(str, 10, 64)
			if err != nil {
				r.JSON(errCode, err)
			}
			return
		}

		page.SinceId, err = parseInt64(qs.Get("since_id"))
		if err != nil {
			return
		}

		page.MaxId, err = parseInt64(qs.Get("max_id"))
		if err != nil {
			return
		}

		countStr := qs.Get("count")
		if countStr == "" {
			countStr = "0"
		}
		page.Count, err = strconv.Atoi(countStr)
		if err != nil {
			return
		}

		if page.Count > 100 {
			page.Count = 100
		}

		if page.Count == 0 {
			page.Count = 10
		}

		if page.MaxId == 0 {
			page.MaxId = math.MaxInt32
		}

		c.Map(&page)
	}
}
