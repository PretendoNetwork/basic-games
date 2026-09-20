package games

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

var residentEvilRevelations3DS = Game{
	Slug:      "resident-evil-revelations-3ds",
	Name:      "Resident Evil: Revelations (3DS)",
	EnvPrefix: "PN_REVELATIONS",

	AccessKey:                      "f26f5737",
	LibraryVersion:                 nex.NewLibraryVersion(2, 3, 2),
	AuthenticationEndpointStreamID: 1,
	SecureEndpointStreamID:         1,
	SecureEndpointStreamType:       constants.StreamTypeRVSecure,
	TitleIDs:                       []string{"00053B00"},

	Authentication: AuthenticationProtocols{
		TicketGranting: &TicketGranting{},
	},

	Secure: SecureProtocols{
		SecureConnection: &SecureConnection{
			InsecureRegister: true,
			CreateReportDBRecord: func(_ types.PID, _ types.UInt32, _ types.QBuffer) error {
				// * Stub for now, matches the existing repos
				return nil
			},
		},
		NATTraversal:       &NATTraversal{},
		MatchMaking:        &MatchMaking{},
		MatchMakingExt:     &MatchMakingExt{},
		MatchmakeExtension: &MatchmakeExtension{},
	},
}
