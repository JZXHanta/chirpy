package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JZXHanta/chirpy/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits int
	DB             *database.DB
	jwtSecret      string
<<<<<<< HEAD
	polkaKey       string
=======
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
}

func main() {
	const filepathRoot = "."
<<<<<<< HEAD

	godotenv.Load(".env")

	port := os.Getenv("PORT")

=======
	const port = "8080"

	godotenv.Load(".env")

>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
<<<<<<< HEAD

	polkaKey := os.Getenv("POLKA_KEY")
=======
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if dbg != nil && *dbg {
		err := db.ResetDB()
		if err != nil {
			log.Fatal(err)
		}
	}

	apiCfg := apiConfig{
		fileserverHits: 0,
		DB:             db,
		jwtSecret:      jwtSecret,
<<<<<<< HEAD
		polkaKey:       polkaKey,
=======
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
	}

	router := chi.NewRouter()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	router.Handle("/app", fsHandler)
	router.Handle("/app/*", fsHandler)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/reset", apiCfg.handlerReset)

<<<<<<< HEAD
	apiRouter.Post("/polka/webhooks", apiCfg.handlerWebhook)

=======
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
	apiRouter.Post("/revoke", apiCfg.handlerRevoke)
	apiRouter.Post("/refresh", apiCfg.handlerRefresh)
	apiRouter.Post("/login", apiCfg.handlerLogin)

	apiRouter.Post("/users", apiCfg.handlerUsersCreate)
	apiRouter.Put("/users", apiCfg.handlerUsersUpdate)

<<<<<<< HEAD
	apiRouter.Delete("/chirps/{chirpID}", apiCfg.handlerChirpsDelete)
=======
>>>>>>> c7d63917131f0b218b5d9ade923c3033151b47e7
	apiRouter.Post("/chirps", apiCfg.handlerChirpsCreate)
	apiRouter.Get("/chirps", apiCfg.handlerChirpsRetrieve)
	apiRouter.Get("/chirps/{chirpID}", apiCfg.handlerChirpsGet)
	router.Mount("/api", apiRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", apiCfg.handlerMetrics)
	router.Mount("/admin", adminRouter)

	corsMux := middlewareCors(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
