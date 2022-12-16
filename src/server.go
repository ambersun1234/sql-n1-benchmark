package main

import (
	"sql-n1-benchmark/database"
	"sql-n1-benchmark/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDatabaseClient("benchmark")
	if err != nil {
		panic(err)
	}
	db.AutoMigration()

	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()

	ginRouter.POST("init", router.Init(db))
	ginRouter.POST("benchmark", router.Benchmark(db))

	ginRouter.Run()
}
