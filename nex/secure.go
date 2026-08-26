package nex

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"

	"github.com/PretendoNetwork/basic-games/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(globals.Game.SecureEndpointStreamID)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername

	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(globals.Game.LibraryVersion)
	globals.SecureServer.AccessKey = globals.Game.AccessKey

	if globals.Game.DebugRMCLog {
		globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
			request := packet.RMCMessage()
			protocolName, methodName := globals.RMCNames(request.ProtocolID, request.MethodID)

			fmt.Printf("== %s - Secure ==\n", globals.Game.Name)
			fmt.Printf("User: %d\n", packet.Sender().PID())
			fmt.Printf("Protocol: %d (%s)\n", request.ProtocolID, protocolName)
			fmt.Printf("Method: %d (%s)\n", request.MethodID, methodName)
			fmt.Printf("Paramaters: %s\n", hex.EncodeToString(request.Parameters))
			fmt.Println("===============")
		})
	}

	globals.SecureEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Errorf("Secure: %v", err)
	})

	if globals.Game.UsesMatchmaking() {
		globals.MatchmakingManager = common_globals.NewMatchmakingManager(globals.SecureEndpoint, globals.Postgres)
		globals.MatchmakingManager.GetUserFriendPIDs = globals.GetUserFriendPIDs
	}

	registerCommonSecureServerProtocols()

	globals.SecureServer.Listen(globals.SecureServerPort)
}
