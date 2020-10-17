// DailyGakki - handler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)
import tb "gopkg.in/tucnak/telebot.v2"

func startHandler(m *telebot.Message) {
	_ = b.Notify(m.Sender, telebot.Typing)
	// TODO add photos
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n我会每天定是为你发送最可爱的Gakki！")
}

func aboutHandler(m *telebot.Message) {
	_ = b.Notify(m.Sender, telebot.Typing)
	_, _ = b.Send(m.Sender, "欢迎来到每日最可爱的Gakki！\n"+
		"开发者：@BennyThink\n"+
		"Google Photos 地址："+photos)
}

func newHandler(m *telebot.Message) {
	_ = b.Notify(m.Sender, telebot.Typing)

	p := &tb.Photo{File: tb.FromDisk("photos/yui.jpg")}
	_, _ = b.SendAlbum(m.Sender, tb.Album{p})
	_, _ = b.Send(m.Sender, "怎么样，喜欢今日份的Gakki吗🤩")
}

func settingsHandler(m *telebot.Message) {

	_ = b.Notify(m.Sender, telebot.Typing)
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

func registerButtonNextStep(btn tb.Btn, fun func(c *tb.Callback)) {
	log.Infoln("Registering ", btn.Unique)
	b.Handle(&btn, fun)
}
