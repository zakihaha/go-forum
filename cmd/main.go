package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/go-forum/internal/configs"
	"github.com/zakihaha/go-forum/internal/handlers/memberships"
	"github.com/zakihaha/go-forum/internal/handlers/posts"
	membershipRepo "github.com/zakihaha/go-forum/internal/repository/memberships"
	postRepo "github.com/zakihaha/go-forum/internal/repository/posts"
	membershipSvc "github.com/zakihaha/go-forum/internal/service/memberships"
	postSvc "github.com/zakihaha/go-forum/internal/service/posts"
	"github.com/zakihaha/go-forum/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	cfg = configs.GetConfig()
	fmt.Println(cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	postHandler := posts.NewHandler(r, postService)

	postHandler.RegisterRoute()
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
