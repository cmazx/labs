// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/models"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/handlers"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ..\..\gen --name CrudService --spec ..\..\swagger\swagger.yaml --principal interface{}

func configureFlags(api *operations.CrudServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CrudServiceAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL_MODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)
	var IgnoreNotFoundLogger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	var connection *gorm.DB
	var err error
	for {
		connection, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: IgnoreNotFoundLogger,
		})
		if err == nil {
			break
		}
		log.Printf("Connection error:%s\n", err.Error())
		time.Sleep(1 * time.Second)
		log.Println("Reconnecting")
	}
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	api.UserGetHandler = handlers.NewGetHandler(connection)
	api.UserCreateUserHandler = handlers.NewAddHandler(connection)
	api.UserDeleteUserHandler = handlers.NewDeleteHandler(connection)
	api.UserListHandler = handlers.NewListHandler(connection)
	api.UserUpdateUserHandler = handlers.NewPutHandler(connection)
	api.HealthHealthHandler = handlers.NewHealthHandler()

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler

	return handleCORS(handler)
}
