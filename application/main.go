package main

import (
	"antrein/dd-dashboard-auth/application/common/repository"
	"antrein/dd-dashboard-auth/application/common/resource"
	"antrein/dd-dashboard-auth/application/common/usecase"
	"antrein/dd-dashboard-auth/application/grpc"
	"antrein/dd-dashboard-auth/application/rest"
	"antrein/dd-dashboard-auth/model/config"
	"context"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	resource, err := resource.NewCommonResource(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repository.NewCommonRepository(cfg, resource)
	if err != nil {
		log.Fatal(err)
	}
	uc, err := usecase.NewCommonUsecase(cfg, repository)
	if err != nil {
		log.Fatal(err)
	}

	rest_app, err := rest.ApplicationDelegate(cfg, uc, resource)
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server concurrently
	go func() {
		grpc_app, err := grpc.ApplicationDelegate(cfg, uc, resource)
		if err != nil {
			log.Fatal(err)
		}
		if err := grpc.StartServer(cfg, grpc_app); err != nil {
			log.Fatal(err)
		}
	}()

	if err = rest.StartServer(cfg, rest_app); err != nil {
		log.Fatal(err)
	}
}
