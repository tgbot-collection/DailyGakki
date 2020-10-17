// DailyGakki - handler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

func startHandler(m *tb.Message) {
	_ = b.Notify(m.Sender, tb.Typing)
	// TODO add album
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

	var max = 3
	var sendAlbum tb.Album

	chosen := ChoosePhotos(max)
	for _, photoPath := range chosen[1:max] {
		p := &tb.Photo{File: tb.FromDisk(photoPath)}
		sendAlbum = append(sendAlbum, p)
	}
	p := &tb.Photo{File: tb.FromDisk(chosen[0]), Caption: "怎么样，喜欢今日份的Gakki吗🤩"}
	sendAlbum = append(sendAlbum, p)

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

func ChoosePhotos(count int) (paths []string) {
	photoMap := listAll(photos)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= count; i++ {
		index := rand.Intn(len(photoMap))
		paths = append(paths, photoMap[index])
		delete(photoMap, index)
	}

	return
}

func listAll(path string) (photo map[int]string) {
	photo = make(map[int]string)
	files, _ := ioutil.ReadDir(path)
	var start = 0
	for _, fi := range files {
		if !fi.IsDir() {
			photo[start] = filepath.Join(path, fi.Name())
			start += 1
		}
	}
	return
}
