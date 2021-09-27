package chat

const (
	KaiheilaName = "开黑啦"
	KaiheilaID   = -100
	KaiheilaUrl  = "https://kaihei.co/p4Bl4i"
	DiscordName  = "Discord"
	DiscordID    = -101
	DiscordUrl   = "https://discord.gg/nW7N4cy"
	QQName       = "QQ"
	QQID         = -102
	QQUrl        = "https://jq.qq.com/?_wv=1027&k=8khzZZ5s"
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
		return ServerURL
	}
}
