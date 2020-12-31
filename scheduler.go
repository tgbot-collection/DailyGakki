// DailyGakki - scheduler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func scheduler() {
	// TODO
	log.Infoln("Start scheduler...")

	currentWindow := time.Now().Format("15:04")

	allData := readJSON()
	var sendList = make(map[string][]int64)

	for _, u := range allData {
		for _, t := range u.Time {
			sendList[t] = append(sendList[t], u.ChatId)
		}
	}

	log.Infof("Total count: %d", len(sendList))

	for _, v := range sendList[currentWindow] {
		// v is user id
		log.Infof("Send message to: %d", v)

		m := tb.Message{
			Sender: &tb.User{ID: int(v)},
		}
		// If you're sending bulk notifications to multiple users,
		//the API will not allow more than 30 messages per second or so.
		time.Sleep(time.Second * 5)
		_ = b.Notify(m.Sender, tb.Typing)
		sendAlbum := generatePhotos()
		_ = b.Notify(m.Sender, tb.UploadingPhoto)
		_, _ = b.SendAlbum(m.Sender, sendAlbum)
	}

}
