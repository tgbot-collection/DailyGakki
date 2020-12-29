// DailyGakki - config
// 2020-10-17 14:14
// Benny <benny.think@gmail.com>

package main

import "os"

var album = "https://photos.app.goo.gl/2aLeoBiRypWRR8yY9"
var photosPath = os.Getenv("PHOTOS")
var token = os.Getenv("TOKEN")
var reviewer = os.Getenv("REVIEWER")

type User struct {
	ChatId int64 `json:"chat_id"`
	Count  string
	Time   int64
}
