package server

import (
	"context"
	"net"
	"pikpo_exam/config"
	db "pikpo_exam/pkg/db"
	"pikpo_exam/pkg/entities"
	todolog "pikpo_exam/pkg/log"
	"pikpo_exam/pkg/services"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func Run() error {

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := config.Reload(nil)
	if err != nil {
		log.Error().Err(err).Msg("error processing config data")
		return err
	}

	todolog.InitLogging()

	database, dbHandler, err := db.ConnectDB(config.NewDatabaseConfig())
	if err != nil {
		log.Error().Err(err).Msg("error connecting database")
		return err
	}

	models := entities.GetAllModels()
	err = database.AutoMigrate(models...)
	if err != nil {
		log.Error().Err(err).Msg("error run migration")
		return err
	}

	server, lis, err := startGRPCServer(config.GetListenPort())
	if err != nil {
		log.Fatal().Err(err).Msgf("failed starting server: %s", err.Error())
		return err
	}

	services.RegisterService(server, dbHandler)

	err = runGRPCServer(server, lis)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed running server: %s", err.Error())
		return err
	}

	return nil
}

func startGRPCServer(port string) (*grpc.Server, net.Listener, error) {
	log.Debug().Msg("[GRPC Server] starting the server...")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, nil, err
	}

	server := grpc.NewServer()

	return server, lis, nil
}

func runGRPCServer(server *grpc.Server, lis net.Listener) error {
	log.Debug().Msgf("[GRPC Server] try to run the server on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		return err
	}

	log.Debug().Msgf("[GRPC Server] server running!...")
	return nil
}
