package chat

import (
	"strings"
	"unicode"
)

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
	ServerURL    = "https://game.new-page.xyz/play#servers"
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

func RemoveColor(str string) string {
	return removeColor(removeInvalidChar(str))
}

func removeInvalidChar(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, str)
}

func removeColor(str string) string {
	s := strings.ReplaceAll(str, "{normal}", "")
	s = strings.ReplaceAll(s, "{default}", "")
	s = strings.ReplaceAll(s, "{white}", "")
	s = strings.ReplaceAll(s, "{darkred}", "")
	s = strings.ReplaceAll(s, "{pink}", "")
	s = strings.ReplaceAll(s, "{green}", "")
	s = strings.ReplaceAll(s, "{lime}", "")
	s = strings.ReplaceAll(s, "{yellow}", "")
	s = strings.ReplaceAll(s, "{lightgreen}", "")
	s = strings.ReplaceAll(s, "{lightred}", "")
	s = strings.ReplaceAll(s, "{red}", "")
	s = strings.ReplaceAll(s, "{gray}", "")
	s = strings.ReplaceAll(s, "{grey}", "")
	s = strings.ReplaceAll(s, "{olive}", "")
	s = strings.ReplaceAll(s, "{orange}", "")
	s = strings.ReplaceAll(s, "{silver}", "")
	s = strings.ReplaceAll(s, "{lightblue}", "")
	s = strings.ReplaceAll(s, "{blue}", "")
	s = strings.ReplaceAll(s, "{purple}", "")
	s = strings.ReplaceAll(s, "{darkorange}", "")
	s = strings.ReplaceAll(s, "{name}", "")

	return s
}
