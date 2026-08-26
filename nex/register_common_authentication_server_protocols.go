package nex

import (
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	commonticketgranting "github.com/PretendoNetwork/nex-protocols-common-go/v2/ticket-granting"
	ticketgranting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"

	"github.com/PretendoNetwork/basic-games/globals"
)

func registerCommonAuthenticationServerProtocols() {
	if globals.Game.Authentication.TicketGranting != nil {
		ticketGrantingProtocol := ticketgranting.NewProtocol()
		globals.AuthenticationEndpoint.RegisterServiceProtocol(ticketGrantingProtocol)
		commonTicketGrantingProtocol := commonticketgranting.NewCommonProtocol(ticketGrantingProtocol)
		commonTicketGrantingProtocol.ConfigurePNValidation(globals.Game.TitleIDs)

		if globals.Game.Authentication.TicketGranting.InsecureLogin {
			commonTicketGrantingProtocol.EnableInsecureLogin()
		}

		secureStationURL := types.NewStationURL("")
		secureStationURL.SetURLType(constants.StationURLPRUDPS)
		secureStationURL.SetAddress(globals.SecureServerHost)
		secureStationURL.SetPortNumber(uint16(globals.SecureServerPort))
		secureStationURL.SetConnectionID(1)
		secureStationURL.SetPrincipalID(types.NewPID(2))
		secureStationURL.SetStreamID(globals.Game.SecureEndpointStreamID)
		secureStationURL.SetStreamType(globals.Game.SecureEndpointStreamType)
		secureStationURL.SetType(uint8(constants.StationURLFlagPublic))

		commonTicketGrantingProtocol.SecureStationURL = secureStationURL
		commonTicketGrantingProtocol.BuildName = types.NewString(globals.Game.BuildName)
		commonTicketGrantingProtocol.SecureServerAccount = globals.SecureServerAccount
	}
}
