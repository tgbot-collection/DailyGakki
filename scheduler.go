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
	log.Infoln("Start scheduler...")
	sendList := readJSON()
	log.Infof("Total count: %d", len(sendList))

	for _, v := range sendList {
		log.Infof("Send message to: %d", v.ChatId)

		m := tb.Message{
			Sender: &tb.User{ID: v.ChatId},
		}
		time.Sleep(time.Second * 5)
		_ = b.Notify(m.Sender, tb.Typing)
		sendAlbum := generatePhotos()
		_ = b.Notify(m.Sender, tb.UploadingPhoto)
		_, _ = b.SendAlbum(m.Sender, sendAlbum)
	}

}
