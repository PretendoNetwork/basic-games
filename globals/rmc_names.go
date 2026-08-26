package globals

import (
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	messaging "github.com/PretendoNetwork/nex-protocols-go/v2/messaging"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	nintendo_notifications "github.com/PretendoNetwork/nex-protocols-go/v2/nintendo-notifications"
	notifications "github.com/PretendoNetwork/nex-protocols-go/v2/notifications"
	persistent_store "github.com/PretendoNetwork/nex-protocols-go/v2/persistent-store"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	remote_log_device "github.com/PretendoNetwork/nex-protocols-go/v2/remote-log-device"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
)

// TODO - This belongs in nex-protocols-go but I can't be arsed right now
type protocolInfo struct {
	name    string
	methods map[uint32]string
}

var protocolInfoByID = map[uint16]protocolInfo{
	remote_log_device.ProtocolID: {
		name: "RemoteLogDeviceProtocol",
		methods: map[uint32]string{
			1: "Log",
		},
	},
	nat_traversal.ProtocolID: {
		name: "NATTraversalProtocol",
		methods: map[uint32]string{
			1: "RequestProbeInitiation",
			2: "InitiateProbe",
			3: "RequestProbeInitiationExt",
			4: "ReportNATTraversalResult",
			5: "ReportNATProperties",
			6: "GetRelaySignatureKey",
			7: "ReportNATTraversalDetail",
		},
	},
	ticket_granting.ProtocolID: {
		name: "TicketGrantingProtocol",
		methods: map[uint32]string{
			1: "Login/ValidateAndRequestTicket",
			2: "LoginEx/ValidateAndRequestTicketWithCustomData",
			3: "RequestTicket",
			4: "GetPID",
			5: "GetName",
			6: "LoginWithContext/ValidateAndRequestTicketWithParam",
		},
	},
	secure_connection.ProtocolID: {
		name: "SecureConnectionProtocol",
		methods: map[uint32]string{
			1: "Register",
			2: "RequestConnectionData",
			3: "RequestUrls",
			4: "RegisterEx",
			5: "TestConnectivity",
			6: "UpdateURLs",
			7: "ReplaceURL",
			8: "SendReport",
		},
	},
	notifications.ProtocolID: {
		name: "NotificationProtocol",
		methods: map[uint32]string{
			1: "ProcessNotificationEvent",
		},
	},
	match_making.ProtocolID: {
		name: "MatchMakingProtocol",
		methods: map[uint32]string{
			1:  "RegisterGathering",
			2:  "UnregisterGathering",
			3:  "UnregisterGatherings",
			4:  "UpdateGathering",
			5:  "Invite",
			6:  "AcceptInvitation",
			7:  "DeclineInvitation",
			8:  "CancelInvitation",
			9:  "GetInvitationsSent",
			10: "GetInvitationsReceived",
			11: "Participate",
			12: "CancelParticipation",
			13: "GetParticipants",
			14: "AddParticipants",
			15: "GetDetailedParticipants",
			16: "GetParticipantsURLs",
			17: "FindByType",
			18: "FindByDescription",
			19: "FindByDescriptionRegex",
			20: "FindByID",
			21: "FindBySingleID",
			22: "FindByOwner",
			23: "FindByParticipants",
			24: "FindInvitations",
			25: "FindBySQLQuery",
			26: "LaunchSession",
			27: "UpdateSessionURL",
			28: "GetSessionURL",
			29: "GetState",
			30: "SetState",
			31: "ReportStats",
			32: "GetStats",
			33: "DeleteGathering",
			34: "GetPendingDeletions",
			35: "DeleteFromDeletions",
			36: "MigrateGatheringOwnershipV1",
			37: "FindByDescriptionLike",
			38: "RegisterLocalURL",
			39: "RegisterLocalURLs",
			40: "UpdateSessionHostV1",
			41: "GetSessionURLs",
			42: "UpdateSessionHost",
			43: "UpdateGatheringOwnership",
			44: "MigrateGatheringOwnership",
		},
	},
	messaging.ProtocolID: {
		name: "MessagingProtocol",
		methods: map[uint32]string{
			1: "DeliverMessage",
			2: "GetNumberOfMessages",
			3: "GetMessagesHeaders",
			4: "RetrieveAllMessagesWithinRange",
			5: "RetrieveMessages",
			6: "DeleteMessages",
			7: "DeleteAllMessages",
			8: "DeliverMessageMultiTarget",
		},
	},
	persistent_store.ProtocolID: {
		name: "PersistentStoreProtocol",
		methods: map[uint32]string{
			1: "FindByGroup",
			2: "InsertItem",
			3: "RemoveItem",
			4: "GetItem",
			5: "InsertCustomItem",
			6: "GetCustomItem",
			7: "FindItemsBySQLQuery",
		},
	},
	message_delivery.ProtocolID: {
		name: "MessageDeliveryProtocol",
		methods: map[uint32]string{
			1: "DeliverMessage",
			2: "DeliverMessageMultiTarget",
		},
	},
	match_making_ext.ProtocolID: {
		name: "MatchMakingProtocolExt",
		methods: map[uint32]string{
			1: "EndParticipation",
			2: "GetParticipants",
			3: "GetDetailedParticipants",
			4: "GetParticipantsURLs",
			5: "GetGatheringRelations",
			6: "DeleteFromDeletions",
		},
	},
	nintendo_notifications.ProtocolID: {
		name: "NintendoNotificationEventProtocol",
		methods: map[uint32]string{
			1: "ProcessNintendoNotificationEvent",
			2: "ProcessNintendoNotificationEvent",
		},
	},
	matchmake_extension.ProtocolID: {
		name: "MatchmakeExtensionProtocol",
		methods: map[uint32]string{
			1:  "CloseParticipation",
			2:  "OpenParticipation",
			3:  "AutoMatchmake_Postpone",
			4:  "BrowseMatchmakeSession",
			5:  "BrowseMatchmakeSessionWithHostUrls",
			6:  "CreateMatchmakeSession",
			7:  "JoinMatchmakeSession",
			8:  "ModifyCurrentGameAttribute",
			9:  "UpdateNotificationData",
			10: "GetFriendNotificationData",
			11: "UpdateApplicationBuffer",
			12: "UpdateMatchmakeSessionAttribute",
			13: "GetlstFriendNotificationData",
			14: "UpdateMatchmakeSession",
			15: "AutoMatchmakeWithSearchCriteria_Postpone",
			16: "GetPlayingSession",
			17: "CreateCommunity",
			18: "UpdateCommunity",
			19: "JoinCommunity",
			20: "FindCommunityByGatheringId",
			21: "FindOfficialCommunity",
			22: "FindCommunityByParticipant",
			23: "UpdatePrivacySetting",
			24: "GetMyBlackList",
			25: "AddToBlackList",
			26: "RemoveFromBlackList",
			27: "ClearMyBlackList",
			28: "ReportViolation",
			29: "IsViolationUser",
			30: "JoinMatchmakeSessionEx",
			31: "GetSimplePlayingSession",
			32: "GetSimpleCommunity",
			33: "AutoMatchmakeWithGatheringId_Postpone",
			34: "UpdateProgressScore",
			35: "DebugNotifyEvent",
			36: "GenerateMatchmakeSessionSystemPassword",
			37: "ClearMatchmakeSessionSystemPassword",
			38: "CreateMatchmakeSessionWithParam",
			39: "JoinMatchmakeSessionWithParam",
			40: "AutoMatchmakeWithParam_Postpone",
			41: "FindMatchmakeSessionByGatheringIdDetail",
			42: "BrowseMatchmakeSessionNoHolder",
			43: "BrowseMatchmakeSessionWithHostUrlsNoHolder",
			44: "UpdateMatchmakeSessionPart",
			45: "RequestMatchmaking",
			46: "WithdrawMatchmaking",
			47: "WithdrawMatchmakingAll",
			48: "FindMatchmakeSessionByGatheringId",
			49: "FindMatchmakeSessionBySingleGatheringId",
			50: "FindMatchmakeSessionByOwner",
			51: "FindMatchmakeSessionByParticipant",
			52: "BrowseMatchmakeSessionNoHolderNoResultRange",
			53: "BrowseMatchmakeSessionWithHostUrlsNoHolderNoResultRange",
			54: "FindCommunityByOwner",
		},
	},
	ranking.ProtocolID: {
		name: "RankingProtocol",
		methods: map[uint32]string{
			1:  "UploadScore",
			2:  "DeleteScore",
			3:  "DeleteAllScores",
			4:  "UploadCommonData",
			5:  "DeleteCommonData",
			6:  "GetCommonData",
			7:  "ChangeAttributes",
			8:  "ChangeAllAttributes",
			9:  "GetRanking",
			10: "GetApproxOrder",
			11: "GetStats",
			12: "GetRankingByPIDList",
			13: "GetRankingByUniqueIdList",
			14: "GetCachedTopXRanking",
			15: "GetCachedTopXRankings",
		},
	},
}

// RMCNames resolves a protocol and method ID to human readable names for
// logging/debugging purposes. Unknown values return the string "Unknown"
func RMCNames(protocolID uint16, methodID uint32) (string, string) {
	info, exists := protocolInfoByID[protocolID]
	if !exists {
		return "Unknown", "Unknown"
	}

	method, exists := info.methods[methodID]
	if !exists {
		return info.name, "Unknown"
	}

	return info.name, method
}
