package nex

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"

	"github.com/PretendoNetwork/basic-games/globals"
)

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(globals.Game.AuthenticationEndpointStreamID)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername

	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(globals.Game.LibraryVersion)
	globals.AuthenticationServer.AccessKey = globals.Game.AccessKey

	if globals.Game.DebugRMCLog {
		globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
			request := packet.RMCMessage()
			protocolName, methodName := globals.RMCNames(request.ProtocolID, request.MethodID)

			fmt.Printf("== %s - Auth ==\n", globals.Game.Name)
			fmt.Printf("User: %d\n", packet.Sender().PID())
			fmt.Printf("Protocol: %d (%s)\n", request.ProtocolID, protocolName)
			fmt.Printf("Method: %d (%s)\n", request.MethodID, methodName)
			fmt.Printf("Paramaters: %s\n", hex.EncodeToString(request.Parameters))
			fmt.Println("===============")
		})
	}

	globals.AuthenticationEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Errorf("Auth: %v", err)
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(globals.AuthenticationServerPort)
}
