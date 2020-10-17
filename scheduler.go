// DailyGakki - scheduler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func scheduler() {
	sendList := readJSON()
	for _, v := range sendList {
		m := tb.Message{
			Sender: &tb.User{ID: v.ChatId},
		}

		_ = b.Notify(m.Sender, tb.Typing)
		sendAlbum := generatePhotos()
		_ = b.Notify(m.Sender, tb.UploadingPhoto)
		_, _ = b.SendAlbum(m.Sender, sendAlbum)
	}

}
