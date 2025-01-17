package server

import (
	"net/http"

	"github.com/gabyross/superMercado/internal/handler"
	"github.com/gabyross/superMercado/internal/repository"
	"github.com/gabyross/superMercado/internal/service"
	"github.com/go-chi/chi"
)

type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(config *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "./../../products.json",
	}
	if config != nil {
		if config.ServerAddress != "" {
			defaultConfig.ServerAddress = config.ServerAddress
		}
		if config.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = config.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (server *ServerChi) Run() (err error) {
	// - repository
	pr, err := repository.NewProductRepository(server.loaderFilePath)
	if err != nil {
		return
	}
	// - service
	ps := service.NewProductService(pr)

	// - handler
	ph := handler.NewProductHandler(ps)
	ah := handler.NewAliveHandler()

	// router
	router := createRouter(ah, ph)

	// run server
	err = http.ListenAndServe(server.serverAddress, router)

	return
}

func createRouter(ah *handler.AliveHandler, ph *handler.ProductHandler) *chi.Mux {
	router := chi.NewRouter()

	// Health check endpoint
	router.Get("/ping", ah.Alive)

	// Products endpoints
	router.Route("/products", func(r chi.Router) {
		r.Get("/", ph.GetAllProducts)
		r.Get("/{id}", ph.GetProductByID)
		r.Get("/search", ph.SearchProductByPriceGreaterThan)
		r.Post("/", ph.CreateProduct)
		r.Put("/{id}", ph.UpdateProduct)
		r.Patch("/{id}", ph.PatchProduct)
		r.Delete("/{id}", ph.DeleteProduct)

	})

	return router
}
