package nex

import (
	commonmatchmaking "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	commonmatchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	commonmatchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"

	matchmaking "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	"github.com/PretendoNetwork/basic-games/globals"
)

func registerCommonSecureServerProtocols() {
	if globals.Game.Secure.SecureConnection != nil {
		secureProtocol := secure.NewProtocol()
		globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
		commonSecureProtocol := commonsecure.NewCommonProtocol(secureProtocol)

		if globals.Game.Secure.SecureConnection.InsecureRegister {
			commonSecureProtocol.EnableInsecureRegister()
		}
	}

	if globals.Game.Secure.NATTraversal != nil {
		natTraversalProtocol := nattraversal.NewProtocol()
		globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
		commonnattraversal.NewCommonProtocol(natTraversalProtocol)
	}

	if globals.Game.Secure.MatchMaking != nil {
		matchMakingProtocol := matchmaking.NewProtocol()
		globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
		commonmatchmaking.NewCommonProtocol(matchMakingProtocol).SetManager(globals.MatchmakingManager)
	}

	if globals.Game.Secure.MatchMakingExt != nil {
		matchMakingExtProtocol := matchmakingext.NewProtocol()
		globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
		commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol).SetManager(globals.MatchmakingManager)
	}

	if globals.Game.Secure.MatchmakeExtension != nil {
		matchmakeExtensionProtocol := matchmakeextension.NewProtocol()
		globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
		commonmatchmakeextension.NewCommonProtocol(matchmakeExtensionProtocol).SetManager(globals.MatchmakingManager)
	}
}
