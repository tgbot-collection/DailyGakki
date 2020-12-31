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
	"strconv"
	"strings"
)

import "github.com/tgbot-collection/tgbot_ping"

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
		"GitHub: https://github.com/tgbot-collection/DailyGakki \n" +
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
	// send out push time
	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}
	add := selector.Data("增加推送时间", "AddPush")
	modify := selector.Data("修改推送时间", "ModifyPush")
	btns = append(btns, add, modify)
	selector.Inline(
		selector.Row(btns...),
	)

	_ = b.Notify(m.Chat, tb.Typing)
	pushTimeStr := strings.Join(getPushTime(m.Chat.ID), " ")
	if pushTimeStr == "" {
		message := fmt.Sprintf("哼假粉😕，都没有 /subscribe 还想看！")
		_, _ = b.Send(m.Chat, message)
	} else {
		message := fmt.Sprintf("你目前的推送时间有：%s，你想要增加还是删除？", pushTimeStr)
		_, _ = b.Send(m.Chat, message, selector)
	}

	///////

	//var inlineKeys [][]tb.InlineButton
	//
	//var unique []tb.InlineButton
	//unique = append(unique, tb.InlineButton{
	//	Unique: fmt.Sprintf("SubTime%s", "18:11"),
	//	Text:   "18:11",
	//})
	//inlineKeys = append(inlineKeys, unique)
	//
	//var btns []tb.InlineButton
	//var count = 1
	//for _, t := range timeSeries() {
	//	if count <= 5 {
	//		var temp = tb.InlineButton{
	//			Unique: fmt.Sprintf("SubTime%s", t),
	//			Text:   t,
	//		}
	//		btns = append(btns, temp)
	//		count++
	//	} else {
	//		count = 1
	//		inlineKeys = append(inlineKeys, btns)
	//		btns =[]tb.InlineButton{}
	//	}
	//}
	//
	//_, _ = b.Send(m.Sender, "好的，那你选个时间吧！", &tb.ReplyMarkup{
	//	InlineKeyboard: inlineKeys,
	//})

	//////

}

func channelHandler(m *tb.Message) {
	if m.Text == "/subscribe" {
		subHandler(m)
	} else if m.Text == "/unsubscribe" {
		unsubHandler(m)
	} else {
		log.Infof("Oops. %s is not a command. Ignore it.", m.Text)
	}
}

func subHandler(m *tb.Message) {
	// check permission first
	canSubscribe := checkSubscribePermission(m)
	if !canSubscribe {
		log.Infof("Denied subscribe request for: %d", m.Sender.ID)
		_ = b.Notify(m.Chat, tb.Typing)
		_, _ = b.Send(m.Chat, "ええ😉只有管理员才能进行设置哦")
		return
	}

	caption := "已经订阅成功啦！将在每晚18:11准时为你推送最可爱的Gakki！如有需要可在 /settings 中更改时间和频率"
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

	addInitSub(m.Chat.ID)

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
		return
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
	info := tgbot_ping.GetRuntime("botsrunner_gakki_1", "Gakki Bot", "html")
	_, _ = b.Send(m.Chat, info, &tb.SendOptions{ParseMode: tb.ModeHTML})
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
	// private and channel: allow
	// group: check admin
	if m.Private() || m.Chat.Type == "channel" {
		allow = true
	} else {
		admins, _ := b.AdminsOf(m.Chat)
		for _, admin := range admins {
			if admin.User.ID == m.Sender.ID {
				allow = true
			}
		}
	}
	return
}

func photoHandler(m *tb.Message) {
	if !m.Private() {
		return
	}
	userID, _ := strconv.Atoi(reviewer)
	mm := tb.Message{
		Sender: &tb.User{
			ID: userID,
		},
	}

	_ = b.Notify(m.Chat, tb.Typing)
	botSent, _ := b.Reply(m, "你的Review已经发出去惹……请耐心等待😄")

	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}
	p1, p2 := botSent.MessageSig()
	data := fmt.Sprintf("%v|%v", p1, p2)
	approve := selector.Data("Yes", "Yes", data)
	deny := selector.Data("No", "No", data)

	btns = append(btns, approve, deny)

	selector.Inline(
		selector.Row(btns...),
	)

	fwd, err := b.Forward(mm.Sender, m, selector)
	if err != nil {
		log.Errorln(err)
		_, _ = b.Edit(botSent, "呃……由于某种神秘的原因，Review请求发送失败了，你再发一下试试\n"+err.Error())
	} else {
		_, _ = b.Reply(fwd, "请Review", selector)

	}

}

func callbackEntrance(c *tb.Callback) {
	log.Infof("Initiating callback from %d", c.Sender.ID)
	// this callback interacts with requester
	switch {
	case c.Data == "\fAddPush":
	case strings.HasPrefix(c.Data, "\fYes"):
		approveCallback(c)
	case strings.HasPrefix(c.Data, "\fNo"):
		denyCallback(c)
	}
}

func getStoredMessage(data string) tb.StoredMessage {
	// data Yes|5159|123456789
	splits := strings.Split(data, "|")
	cid, _ := strconv.ParseInt(splits[2], 10, 64)
	botM := tb.StoredMessage{MessageID: splits[1], ChatID: cid}
	return botM
}

func approveCallback(c *tb.Callback) {
	log.Infof("approve new photos from %s", c.Data)
	botM := getStoredMessage(c.Data)

	approveAction(c.Message.ReplyTo)
	_ = b.Respond(c, &tb.CallbackResponse{Text: "Approved"})
	_, _ = b.Edit(botM, "你的图片被接受了😊")

	_ = b.Delete(c.Message)         // this message
	_ = b.Delete(c.Message.ReplyTo) // original message with photo
}

func denyCallback(c *tb.Callback) {
	log.Infof("deny new photos from %s", c.Data)
	botM := getStoredMessage(c.Data)

	_ = b.Respond(c, &tb.CallbackResponse{Text: "Denied"})
	_, _ = b.Edit(botM, "你的图片被拒绝了😫")

	_ = b.Delete(c.Message)         // this message
	_ = b.Delete(c.Message.ReplyTo) // original message with photo
}

func approveAction(reviewMessage *tb.Message) {
	// this handler interacts with reviewer
	photo := reviewMessage.Photo
	document := reviewMessage.Document
	var filename = ""
	var fileobject tb.File

	if photo != nil {
		filename = photo.UniqueID + ".jpg"
		fileobject = photo.File
	} else if document != nil {
		filename = document.UniqueID + ".jpg"
		fileobject = document.File
	} else {
		return
	}
	picPath := filepath.Join(photosPath, filename)
	log.Infof("Downloading photos to %s", picPath)
	err = b.Download(&fileobject, picPath)
	if err != nil {
		log.Errorln("Download failed", err)
	}
}

func submitHandler(m *tb.Message) {

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "想要向我提交新的图片吗？直接把图片发送给我就可以！单张，多张为一组，转发都可以的！\n"+
		"文件和图片的形式发送给bot都可以哦。如有问题可以联系 @BennyThink")

}

func inline(q *tb.Query) {
	var urls []string
	var web = "https://bot.gakki.photos/"

	for _, p := range ChoosePhotos(3) {
		urls = append(urls, web+filepath.Base(p))
	}

	results := make(tb.Results, len(urls)) // []tb.Result
	for i, url := range urls {
		results[i] = &tb.PhotoResult{
			URL:      url,
			ThumbURL: url,
		}
		// needed to set a unique string ID for each result
		results[i].SetResultID(strconv.Itoa(i))
	}

	log.Infof("Inline pic %v", urls)
	err := b.Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 60, // a minute
	})

	if err != nil {
		log.Println(err)
	}
}
