package main

import (
	"github.com/New-pal/np-backend/auth"
	"github.com/New-pal/np-backend/category"
	"github.com/New-pal/np-backend/common/config"
	"github.com/New-pal/np-backend/common/db"
	"github.com/New-pal/np-backend/core"
	"github.com/New-pal/np-backend/docs"
	"github.com/New-pal/np-backend/user"
	gwt "github.com/ennaque/go-gin-jwt"
	gwtstorage "github.com/ennaque/go-gin-jwt/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"time"
)

func main() {
	config.Setup()
	initSwagger()

	database := db.NewConnection().DB
	sqlDb, _ := database.DB()
	defer sqlDb.Close()
	migrate(database)

	router := gin.Default()

	bindRoutes(router, database)

	if err := router.Run(":9099"); err != nil {
		panic(err)
	}
}

func initJwt(db *gorm.DB, ur *user.UserRepository) *gwt.Gwt {
	st, err := gwtstorage.InitGormStorage(db, viper.GetString("auth.table_prefix"))
	if err != nil {
		panic(err)
	}

	jwt, gwtErr := gwt.Init(gwt.Settings{
		AccessSecretKey: []byte(viper.GetString("auth.access_secret_key")),
		Storage:         st,
		AccessLifetime:  time.Minute * viper.GetDuration("auth.access_lifetime_minutes"),
		RefreshLifetime: time.Hour * viper.GetDuration("auth.refresh_lifetime_hours"),
		SigningMethod:   viper.GetString("auth.signing_method"),
		AuthHeadName:    viper.GetString("auth.auth_head_name"),
		Authenticator:   auth.Authenticator(ur),
		GetUserFunc: func(userId string) (interface{}, error) {
			return ur.GetUserById(userId)
		},
	})
	if gwtErr != nil {
		panic(gwtErr)
	}
	return jwt
}

func bindRoutes(router *gin.Engine, db *gorm.DB) {
	router.Use(core.CorsMiddleware(), core.ContentTypeMiddleware())

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(db)
	userSettingsRepository := user.NewUserSettingsRepository(db)
	userSettingsService := user.NewUserSettingsService(db)
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(db)
	subcategoryRepository := category.NewSubcategoryRepository(db)
	subcategoryService := category.NewSubcategoryService(db)

	jwt := initJwt(db, userRepository)

	apiGroup := router.Group("/api")
	apiGroup.Use(jwt.Middleware.GetAuthMiddleware())
	authGroup := router.Group("/auth")

	auth.BindRouter(authGroup, auth.NewAuthHandler(userRepository, userService, jwt.Handler))
	user.BindRouter(apiGroup, user.NewUserHandler(userRepository, userService, userSettingsRepository,
		userSettingsService))
	category.BindRouter(apiGroup, category.NewCategoryHandler(categoryRepository, categoryService,
		subcategoryRepository, subcategoryService))

	if !viper.GetBool("is_prod") {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func migrate(database *gorm.DB) {
	user.Migrate(database)
	category.Migrate(database)
}

func initSwagger() {
	docs.SwaggerInfo.Title = "Newpal API"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "Backend for newpal application"
	docs.SwaggerInfo.Host = "localhost:9099"
}
