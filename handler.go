// DailyGakki - handler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"path/filepath"
)

func startHandler(m *tb.Message) {
	caption := "欢迎来到每日最可爱的Gakki！\n我会每天定是为你发送最可爱的Gakki！"
	filename := "start.gif"

	log.Infof("Start command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

}

func aboutHandler(m *tb.Message) {
	caption := "欢迎来到每日最可爱的Gakki！\n" +
		"开发者：@BennyThink\n" +
		"GitHub: https://github.com/BennyThink/DailyGakki/" +
		"Google Photos 地址：" + album
	filename := "about.gif"

	log.Infof("About command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

}

func newHandler(m *tb.Message) {
	log.Infof("New command: %d", m.Chat.ID)

	// 默认发送3张
	_ = b.Notify(m.Chat, tb.Typing)
	sendAlbum := generatePhotos()
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	_, _ = b.SendAlbum(m.Chat, sendAlbum)

}

func settingsHandler(m *tb.Message) {
	log.Infof("Settings command: %d", m.Chat.ID)

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "在这里可以设置每日推送时间和每日推送次数")
	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}

	btn := selector.Data("Placeholder", fmt.Sprintf("Placeholder%s%d",
		"Placeholder", m.Chat.ID), "Placeholder")
	//registerButtonNextStep(btn, "addServiceButton")
	btns = append(btns, btn)

	selector.Inline(
		selector.Row(btns...),
	)

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "呀！这部分功能还没做！😅", selector)

}

//func registerButtonNextStep(btn tb.Btn, fun func(c *tb.Callback)) {
//	log.Infoln("Registering ", btn.Unique)
//	b.Handle(&btn, fun)
//}

func subHandler(m *tb.Message) {
	// check permission first
	canSubscribe := checkSubscribePermission(m)
	if !canSubscribe {
		log.Infof("Denied subscribe request for: %d", m.Sender.ID)
		_ = b.Notify(m.Chat, tb.Typing)
		_, _ = b.Send(m.Chat, "ええ😉只有管理员才能进行设置哦")
		return
	}

	caption := "已经订阅成功啦！将在每晚18:11准时为你推送最可爱的Gakki！"
	filename := "sub.gif"

	log.Infof("Sub command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

	add(m.Chat.ID)

}

func unsubHandler(m *tb.Message) {
	canSubscribe := checkSubscribePermission(m)
	if !canSubscribe {
		log.Infof("Denied subscribe request for: %d", m.Sender.ID)
		_ = b.Notify(m.Chat, tb.Typing)
		_, _ = b.Send(m.Chat, "ええ😉只有管理员才能进行设置哦")
		return
	}
	caption := "Gakki含泪挥手告别😭"
	filename := "unsub.gif"

	log.Infof("Unsub command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "😭")

	// 读取文件，增加对象，然后写入
	remove(m.Chat.ID)

}

func messageHandler(m *tb.Message) {
	caption := "私は　今でも空と恋をしています。"
	var filename string

	log.Infof("Message Handler: %d", m.Chat.ID)

	switch m.Text {
	case "😘":
		filename = "kiss.gif"
	case "😚":
		filename = "kiss.gif"
	case "😗":
		filename = "kiss.gif"
	case "❤️":
		filename = "heart1.gif"
	case "❤️❤️":
		filename = "heart2.gif"
	case "❤️❤️❤️":
		filename = "heart3.gif"
	case "🌹":
		filename = "rose.gif"
	case "🦎":
		filename = "lizard.gif"
	default:
		filename = "default.gif"
	}

	log.Infof("Choose %s for text %s", filename, m.Text)
	data, _ := Asset(filepath.Join("images", filename))

	log.Infof("Send %s now...", filename)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)

	}

}

func pingHandler(m *tb.Message) {
	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "pong")
}

func statusHandler(m *tb.Message) {
	_ = b.Notify(m.Chat, tb.Typing)
	currentJSON := readJSON()
	var isSub = false
	for _, user := range currentJSON {
		if user.ChatId == m.Chat.ID {
			isSub = true
		}
	}
	if isSub {
		_, _ = b.Send(m.Chat, "Gakki与你同在😄")
	} else {
		_, _ = b.Send(m.Chat, "还木有每日Gakki💔")

	}
}

func checkSubscribePermission(m *tb.Message) (allow bool) {
	allow = false
	if !m.Private() {
		admins, _ := b.AdminsOf(m.Chat)
		for _, admin := range admins {
			if admin.User.ID == m.Sender.ID {
				allow = true
			}
		}
	} else {
		allow = true
	}
	return
}
