package cmd

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/appcontext"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/logger"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/repository"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/server"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/service"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "ecommerce-server",
	Long:  `Ecommerce Server is a service that provides ecommerce features to users.`,
	Short: `Ecommerce Server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config()
		log := logger.NewLogger(nil)
		app := appcontext.NewAppContext(cfg)
		var err error
		log.Info("Starting the server")
		var dbMysql *gorm.DB

		dbMysql, err = app.GetDBInstance()
		if err != nil {
			log.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
			return
		}
		autoMigrate := cmd.Flag("auto-migrate")
		if autoMigrate == nil || (autoMigrate != nil && autoMigrate.Value.String() == "false") {
			log.Info("Auto migrate")
			database.Migrate(dbMysql)
		}
		cachePool, err := app.GetCachedPool()
		if err != nil {
			log.Fatalf("Failed to start, error connect to DB Redis | %v", err)
			return
		}

		opt := commons.Options{
			Config:    cfg,
			DbMysql:   dbMysql,
			Logger:    log,
			CachePool: cachePool,
		}

		start(opt)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", ".env", "config file (default is $PWD/.env)")
	rootCmd.PersistentFlags().BoolP("auto", "a", true, "automatic running migration (default is true)")
	cobra.OnInitialize()
}

func start(opt commons.Options) {

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	service := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	server := server.NewServer(opt, service)

	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	cacheRepo := repository.NewCacheRepository(repoOption)
	userRepo := repository.NewUserRepository(repoOption)
	authRepo := repository.NewAuthRepository(repoOption)
	repo := repository.Repository{
		Auth:  authRepo,
		Cache: cacheRepo,
		User:  userRepo,
	}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	hc := service.NewHealthCheck(serviceOption)
	user := service.NewUserService(serviceOption)
	auth := service.NewAuthService(serviceOption)
	svc := service.Services{
		HealthCheck: hc,
		Auth:        auth,
		User:        user,
	}

	return &svc
}
