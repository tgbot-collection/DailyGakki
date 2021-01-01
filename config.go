// DailyGakki - config
// 2020-10-17 14:14
// Benny <benny.think@gmail.com>

package main

import "os"

var album = "https://photos.app.goo.gl/2aLeoBiRypWRR8yY9"
var photosPath = os.Getenv("PHOTOS")
var token = os.Getenv("TOKEN")
var reviewer = os.Getenv("REVIEWER")

type userConfig struct {
	ChatId int64    `json:"chat_id"`
	Time   []string `json:"time"`
}

type user map[int64]*userConfig
