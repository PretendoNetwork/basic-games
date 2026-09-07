package games

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

var tetrisAxis = Game{
	Slug:      "tetris-axis",
	Name:      "Tetris Axis",
	EnvPrefix: "PN_TETR",

	AccessKey:                      "2ef57176",
	LibraryVersion:                 nex.NewLibraryVersion(2, 0, 0),
	AuthenticationEndpointStreamID: 1,
	SecureEndpointStreamID:         1,
	SecureEndpointStreamType:       constants.StreamTypeRVSecure,
	BuildName:                      "branch:trunk build:2_15_7221_0",
	TitleIDs:                       []string{"00039E00"},

	Authentication: AuthenticationProtocols{
		TicketGranting: &TicketGranting{},
	},

	Secure: SecureProtocols{
		SecureConnection: &SecureConnection{
			InsecureRegister: true,
			CreateReportDBRecord: func(_ types.PID, _ types.UInt32, _ types.QBuffer) error {
				// * Stub for now, matches the existing repo
				return nil
			},
		},
		NATTraversal:       &NATTraversal{},
		MatchMaking:        &MatchMaking{},
		MatchMakingExt:     &MatchMakingExt{},
		MatchmakeExtension: &MatchmakeExtension{},
	},
}
