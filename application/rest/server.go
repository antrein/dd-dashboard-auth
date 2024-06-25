package rest

import (
	"antrein/dd-dashboard-auth/model/config"
	"fmt"
	"log"
	"net/http"
)

func StartServer(cfg *config.Config, handler http.Handler) error {
	port := cfg.Server.Rest.Port
	address := fmt.Sprintf(":%s", port)

	fmt.Printf("REST app is starting on http://localhost:%s/dd/dashboard/auth\n", port)

	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}
	return nil
}
