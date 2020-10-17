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
	log.Infof("Start command: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	filename := "start.gif"
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename}
	_, _ = b.Send(m.Sender, p)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n我会每天定是为你发送最可爱的Gakki！")

}

func aboutHandler(m *tb.Message) {
	log.Infof("About command: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	filename := "about.gif"
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename}
	_, _ = b.Send(m.Sender, p)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n"+
		"开发者：@BennyThink\n"+
		"GitHub: https://github.com/BennyThink/DailyGakki/"+
		"Google Photos 地址："+album)

}

func newHandler(m *tb.Message) {
	log.Infof("New command: %d", m.Sender.ID)

	// 默认发送3张
	_ = b.Notify(m.Sender, tb.Typing)
	sendAlbum := generatePhotos()
	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	_, _ = b.SendAlbum(m.Sender, sendAlbum)

}

func settingsHandler(m *tb.Message) {
	log.Infof("Settings command: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "在这里可以设置每日推送时间和每日推送次数")
	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}

	btn := selector.Data("Placeholder", fmt.Sprintf("Placeholder%s%d", "Placeholder", m.Sender.ID), "Placeholder")
	//registerButtonNextStep(btn, "addServiceButton")
	btns = append(btns, btn)

	selector.Inline(
		selector.Row(btns...),
	)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "呀！这部分功能还没做！😅", selector)

}

//func registerButtonNextStep(btn tb.Btn, fun func(c *tb.Callback)) {
//	log.Infoln("Registering ", btn.Unique)
//	b.Handle(&btn, fun)
//}

func subHandler(m *tb.Message) {
	log.Infof("Sub command: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	filename := "sub.gif"
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename}
	_, _ = b.Send(m.Sender, p)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "已经订阅成功啦！将在每晚18:11准时为你推送最可爱的Gakki！")
	// 读取文件，增加对象，然后写入
	var this = User{
		ChatId: m.Sender.ID,
		Count:  "",
		Time:   0,
	}
	currentDB := readJSON()
	add(currentDB, this)

}

func unsubHandler(m *tb.Message) {
	log.Infof("Unsub command: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	filename := "unsub.gif"
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename}
	_, _ = b.Send(m.Sender, p)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "Gakki含泪挥手告别😭")
	// 读取文件，增加对象，然后写入

	var this = User{
		ChatId: m.Sender.ID,
		Count:  "",
		Time:   0,
	}
	currentDB := readJSON()
	remove(currentDB, this)

}

func messageHandler(m *tb.Message) {
	log.Infof("Message Handler: %d", m.Sender.ID)

	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "私は　今でも空と恋をしています。")

	var filename string
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
	data, err := Asset(filepath.Join("images", filename))
	if err != nil {
		log.Warningf("File not found %v", err)
	} else {
		log.Infof("Send %s now...", filename)
		_ = b.Notify(m.Sender, tb.UploadingPhoto)
		p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename}
		_, _ = b.Send(m.Sender, p)
	}

}
