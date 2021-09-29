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
	s := strings.Replace(str, "{normal}", "", -1)
	s = strings.Replace(s, "{default}", "", -1)
	s = strings.Replace(s, "{white}", "", -1)
	s = strings.Replace(s, "{darkred}", "", -1)
	s = strings.Replace(s, "{pink}", "", -1)
	s = strings.Replace(s, "{green}", "", -1)
	s = strings.Replace(s, "{lime}", "", -1)
	s = strings.Replace(s, "{yellow}", "", -1)
	s = strings.Replace(s, "{lightgreen}", "", -1)
	s = strings.Replace(s, "{lightred}", "", -1)
	s = strings.Replace(s, "{red}", "", -1)
	s = strings.Replace(s, "{gray}", "", -1)
	s = strings.Replace(s, "{grey}", "", -1)
	s = strings.Replace(s, "{olive}", "", -1)
	s = strings.Replace(s, "{orange}", "", -1)
	s = strings.Replace(s, "{silver}", "", -1)
	s = strings.Replace(s, "{lightblue}", "", -1)
	s = strings.Replace(s, "{blue}", "", -1)
	s = strings.Replace(s, "{purple}", "", -1)
	s = strings.Replace(s, "{darkorange}", "", -1)
	s = strings.Replace(s, "{name}", "", -1)

	return s
}
