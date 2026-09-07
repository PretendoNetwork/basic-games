// Package games holds the configuration data for each basic game
package games

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Game stores the per-game configuration data for a given game
type Game struct {
	// Slug identifies the game via the PN_GAME environment variable
	Slug string

	// Name is the human readable name for the game server
	Name string

	// EnvPrefix is the games environment variable prefix used ot load certain configurable data
	// at runtime. For example the games secure server port is loaded from "PREFIX_SECURE_SERVER_PORT",
	// so setting "PN_TETR" creates "PN_TETR_SECURE_SERVER_PORT"
	EnvPrefix string

	// AccessKey is the games PRUDP access key
	AccessKey string

	// LibraryVersion is the NEX library version the game was built against
	LibraryVersion *nex.LibraryVersion

	// AuthenticationEndpointStreamID is the PRUDP virtual port stream ID that the authentication
	// endpoint binds to
	AuthenticationEndpointStreamID uint8

	// TODO - Also create AuthenticationEndpointStreamType?

	// SecureEndpointStreamID is the PRUDP virtual port stream ID that the secure endpoint binds to
	SecureEndpointStreamID uint8

	// SecureEndpointStreamID is the PRUDP virtual port stream type that the secure endpoint binds to
	SecureEndpointStreamType constants.StreamType

	// BuildName is the NEX server build string for the game server
	BuildName string

	// TitleIDs are the title IDs tied to this game server. Used to validate incoming NEX tokens
	TitleIDs []string

	// DebugRMCLog enables the logging of RMC requests coming from clients. These
	// logs print the user PID, protocol name/method, and the RMC paramaters to the
	// console
	DebugRMCLog bool

	// Authentication holds the protocols on the authentication endpoint
	Authentication AuthenticationProtocols

	// Secure holds the protocols on the secure endpoint
	Secure SecureProtocols
}

// AuthenticationProtocols are the protocols on the authentication endpoint.
// A nil field means the protocol is not registered
type AuthenticationProtocols struct {
	TicketGranting *TicketGranting
}

// SecureProtocols are the protocols on the secure endpoint.
// A nil field means the protocol is not registered
type SecureProtocols struct {
	NATTraversal       *NATTraversal
	SecureConnection   *SecureConnection
	MatchMaking        *MatchMaking
	MatchMakingExt     *MatchMakingExt
	MatchmakeExtension *MatchmakeExtension
}

// TicketGranting configures the TicketGrantingProtocol
type TicketGranting struct {
	// InsecureLogin enables the less secure legacy Login method
	// TODO - Remove both InsecureLogin and InsecureRegister, there are NO cases of games using both, a game MUST use one of the secure methods, so we can validate connections there
	InsecureLogin bool
}

// NATTraversal configures the NATTraversalProtocol
type NATTraversal struct{}

// SecureConnection configures the SecureConnectionProtocol
type SecureConnection struct {
	// InsecureRegister enables the less secure legacy Register method
	// TODO - Remove both InsecureLogin and InsecureRegister, there are NO cases of games using both, a game MUST use one of the secure methods, so we can validate connections there
	InsecureRegister bool

	// CreateReportDBRecord sets the handler for the CreateReportDBRecord function in the
	// common SecureConnectionProtocol implementation
	// TODO - This is not great, ideally nex-protocols-common-go should just handle this and not REQUIRE the developer to define this. Should be an override in common and not exposed/required here
	CreateReportDBRecord func(types.PID, types.UInt32, types.QBuffer) error
}

// MatchMaking configures the MatchMakingProtocol
type MatchMaking struct{}

// MatchMakingExt configures the MatchMakingProtocolExt
type MatchMakingExt struct{}

// MatchmakeExtension configures the MatchmakeExtensionProtocol
type MatchmakeExtension struct{}

// UsesMatchmaking reports whether or not a game uses any of the matchmaking protocols
func (g *Game) UsesMatchmaking() bool {
	return g.Secure.MatchMaking != nil || g.Secure.MatchMakingExt != nil || g.Secure.MatchmakeExtension != nil
}
