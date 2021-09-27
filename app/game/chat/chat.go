package chat

const (
	KaiheilaName = "开黑啦"
	KaiheilaID   = -100
	KaiheilaUrl  = ""
	DiscordName  = "Discord"
	DiscordID    = -101
	DiscordUrl   = ""
	QQName       = "QQ"
	QQID         = -102
	QQUrl        = ""
	ServerURL    = "https://game.new-page.xyz"
)

func GetUrl(serverID int) string {
	switch serverID {
	case KaiheilaID:
		return KaiheilaUrl
	case DiscordID:
		return DiscordUrl
	case QQID:
		return QQUrl
	default:
		return ""
	}
}
