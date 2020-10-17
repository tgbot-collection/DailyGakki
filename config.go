// DailyGakki - config
// 2020-10-17 14:14
// Benny <benny.think@gmail.com>

package main

import "os"

var album = "https://album.app.goo.gl/2aLeoBiRypWRR8yY9"
var photos = os.Getenv("PHOTOS")
var token = os.Getenv("TOKEN")

type User struct {
	ChatId int `json:"chat_id"`
	Count  string
	Time   int64
}
