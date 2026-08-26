package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"os"

	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pbfriends "github.com/PretendoNetwork/grpc/go/friends"

	"github.com/PretendoNetwork/basic-games/games"
	"github.com/PretendoNetwork/basic-games/globals"

	_ "github.com/lib/pq"
)

func init() {
	globals.Logger = plogger.NewLogger()

	var err error

	err = godotenv.Load()
	if err != nil {
		globals.Logger.Warning("Error loading .env file")
	}

	slug := os.Getenv("PN_GAME")
	if slug == "" {
		globals.Logger.Errorf("PN_GAME environment variable not set. Known games: %v", games.Slugs())
		os.Exit(0)
	}

	globals.Game, err = games.Get(slug)
	if err != nil {
		globals.Logger.Error(err.Error())
		os.Exit(0)
	}

	e := &env{
		prefix: globals.Game.EnvPrefix,
	}

	globals.AuthenticationServerPort = e.requiredPort("AUTHENTICATION_SERVER_PORT")
	globals.SecureServerHost = e.requiredString("SECURE_SERVER_HOST")
	globals.SecureServerPort = e.requiredPort("SECURE_SERVER_PORT")

	accountGRPCHost := e.requiredString("ACCOUNT_GRPC_HOST")
	accountGRPCPort := e.requiredPort("ACCOUNT_GRPC_PORT")
	accountGRPCAPIKey := e.lookup("ACCOUNT_GRPC_API_KEY")

	friendsGRPCHost := e.requiredString("FRIENDS_GRPC_HOST")
	friendsGRPCPort := e.requiredPort("FRIENDS_GRPC_PORT")
	friendsGRPCAPIKey := e.lookup("FRIENDS_GRPC_API_KEY")

	postgresURI := e.requiredString("POSTGRES_URI")

	healthCheckPort := e.optionalPort("HEALTH_CHECK_PORT")

	if err := e.err(); err != nil {
		globals.Logger.Errorf("%s: %v", globals.Game.Name, err)
		os.Exit(0)
	}

	if accountGRPCAPIKey == "" {
		globals.Logger.Warning("Insecure gRPC server detected. Account gRPC API key not set")
	}

	kerberosPassword := make([]byte, 0x10)
	_, err = rand.Read(kerberosPassword)
	if err != nil {
		globals.Logger.Error("Error generating Kerberos password")
		os.Exit(0)
	}

	globals.KerberosPassword = string(kerberosPassword)

	globals.InitAccounts()

	common_globals.ConnectToAccountGRPC(accountGRPCHost, uint16(accountGRPCPort), accountGRPCAPIKey)

	if friendsGRPCAPIKey == "" {
		globals.Logger.Warning("Insecure gRPC server detected. Friends gRPC API key not set")
	}

	globals.GRPCFriendsClientConnection, err = grpc.NewClient(fmt.Sprintf("%s:%d", friendsGRPCHost, friendsGRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		globals.Logger.Criticalf("Failed to connect to friends gRPC server: %v", err)
		os.Exit(0)
	}

	globals.GRPCFriendsClient = pbfriends.NewFriendsClient(globals.GRPCFriendsClientConnection)
	globals.GRPCFriendsCommonMetadata = metadata.Pairs(
		"X-API-Key", friendsGRPCAPIKey,
	)

	globals.Postgres, err = sql.Open("postgres", postgresURI)
	if err != nil {
		globals.Logger.Critical(err.Error())
		os.Exit(0)
	}

	if healthCheckPort == 0 {
		globals.Logger.Warning("Basic UDP health check will not be enabled. Health check port not set")
	} else {
		nex.EnableBasicUDPHealthCheck(healthCheckPort)
	}
}
