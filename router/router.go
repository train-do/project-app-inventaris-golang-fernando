package router

import (
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/train-do/project-app-inventaris-golang-fernando/database"
	handler "github.com/train-do/project-app-inventaris-golang-fernando/handler/api"
	"github.com/train-do/project-app-inventaris-golang-fernando/repository"
	"github.com/train-do/project-app-inventaris-golang-fernando/service"
)

func RouterAPI() *chi.Mux {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	handlerGoods := handler.NewGoodsHandler(service.NewGoodsService(repository.NewGoodsRepository(db)))
	handlerCategories := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(db)))
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Route("/items", func(r chi.Router) {
			r.Get("/", handlerGoods.GetAllGoods)
			r.Post("/", handlerGoods.CreateGoods)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlerGoods.GetGoods)
				r.Put("/", handlerGoods.UpdateGoods)
				r.Delete("/", handlerGoods.DeleteGoods)
			})
			r.Route("/investment", func(r chi.Router) {
				r.Get("/", handlerGoods.GetAllInvestment)
				r.Get("/{id}", handlerGoods.GetInvestmentById)
			})
			r.Get("/replacement-needed", handlerGoods.GetReplacementNeeded)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Get("/", handlerCategories.GetAllCategory)
			r.Post("/", handlerCategories.CreateCategory)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlerCategories.GetCategory)
				r.Put("/", handlerCategories.UpdateCategory)
				r.Delete("/", handlerCategories.DeleteCategory)
			})
		})
	})
	// router.Route("/cms", func(r chi.Router) {
	// 	r.Use(middleware.Logger)
	// 	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("Handling GET request"))
	// 	})

	// Grouping /cms/users routes
	// r.Route("/users", func(r chi.Router) {
	// 	r.Get("/", getAllUsers)
	// 	r.Post("/", createUser)
	// 	r.Route("/{userID}", func(r chi.Router) {
	// 		r.Get("/", getUserByID)
	// 		r.Put("/", updateUser)
	// 		r.Delete("/", deleteUser)
	// 	})
	// })

	// // Grouping /cms/products routes
	// r.Route("/products", func(r chi.Router) {
	// 	r.Get("/", getAllProducts)
	// 	r.Post("/", createProduct)
	// 	r.Route("/{productID}", func(r chi.Router) {
	// 		r.Get("/", getProductByID)
	// 		r.Put("/", updateProduct)
	// 		r.Delete("/", deleteProduct)
	// 	})
	// })
	// })
	return router
}
