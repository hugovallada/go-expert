package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/hugovallada/go-expert/api/configs"
	_ "github.com/hugovallada/go-expert/api/docs"
	"github.com/hugovallada/go-expert/api/internal/entity"
	"github.com/hugovallada/go-expert/api/internal/infra/database"
	"github.com/hugovallada/go-expert/api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms

// @contact.name Hugo Lopes
// @contact.url https://github.com/hugovallada
// @contact.email valladahugo@gmail.com

// @license.name Apache 2.0
// @license.url

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productHandler := handlers.NewProductHandler(database.NewProduct(db))
	userHandler := handlers.NewUserHandler(database.NewUser(db))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", cfg.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.Update)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/auth", userHandler.Login)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
