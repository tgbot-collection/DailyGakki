// DailyGakki - handler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func startHandler(m *tb.Message) {
	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n我会每天定是为你发送最可爱的Gakki！")
}

func aboutHandler(m *tb.Message) {
	_ = b.Notify(m.Sender, tb.Typing)
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n"+
		"开发者：@BennyThink\n"+
		"Google Photos 地址："+album)
}

func newHandler(m *tb.Message) {
	// 默认发送3张
	_ = b.Notify(m.Sender, tb.Typing)
	sendAlbum := generatePhotos()
	_ = b.Notify(m.Sender, tb.UploadingPhoto)
	_, _ = b.SendAlbum(m.Sender, sendAlbum)

}

func settingsHandler(m *tb.Message) {
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

func generatePhotos() (sendAlbum tb.Album) {
	var max = 3
	//var sendAlbum tb.Album

	chosen := ChoosePhotos(max)
	for _, photoPath := range chosen[1:max] {
		p := &tb.Photo{File: tb.FromDisk(photoPath)}
		sendAlbum = append(sendAlbum, p)
	}
	p := &tb.Photo{File: tb.FromDisk(chosen[0]), Caption: "怎么样，喜欢今日份的Gakki吗🤩"}
	sendAlbum = append(sendAlbum, p)
	return
}
