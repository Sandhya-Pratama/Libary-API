package router

func SetupRouter(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// Middleware for JWT authentication
	r.Use(middleware.JwtAuthMiddleware(cfg.JwtSecret))

	// Route Handlers
	r.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetBookByID).Methods("GET")
	r.HandleFunc("/api/authors", handler.GetAuthors).Methods("GET")
	r.HandleFunc("/api/categories", handler.GetCategories).Methods("GET")
	r.HandleFunc("/api/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/borrowings", handler.GetBorrowings).Methods("GET")
	r.HandleFunc("/api/borrowings", handler.CreateBorrowing).Methods("POST")
	r.HandleFunc("/api/borrowings/{id}/return", handler.ReturnBorrowing).Methods("POST")

	return r
}
