package router

import (
	"net/http"

	"sql-n1-benchmark/benchmark"
	"sql-n1-benchmark/database"

	"github.com/gin-gonic/gin"
)

type BenchmarkRequest struct {
	Size int `json:"size"`
}

type InitRequest struct {
	Size int `json:"size"`
}

func Benchmark(db *database.DatabaseClient) func(*gin.Context) {
	return func(c *gin.Context) {
		var req BenchmarkRequest

		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := benchmark.Benchmark(db, req.Size); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}

func Init(db *database.DatabaseClient) func(*gin.Context) {
	return func(c *gin.Context) {
		var req InitRequest

		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		db.Init(req.Size)
		c.Status(http.StatusOK)
	}
}
