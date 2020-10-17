// DailyGakki - helper
// 2020-10-17 16:37
// Benny <benny.think@gmail.com>

package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func readJSON() []User {
	log.Infoln("Read json file...")
	jsonFile, _ := os.Open("database.json")
	decoder := json.NewDecoder(jsonFile)

	var db []User
	err = decoder.Decode(&db)
	_ = jsonFile.Close()
	return db

}

func add(d []User, data User) {
	log.Infof("Add subscriber %d", data.ChatId)

	// check and then add
	var shouldWrite = true
	for _, v := range d {
		if v.ChatId == data.ChatId {
			shouldWrite = false
		}
	}
	if shouldWrite {
		d = append(d, data)
		file, _ := json.MarshalIndent(d, "", " ")
		_ = ioutil.WriteFile("database.json", file, 0644)
	}

}

func remove(d []User, data User) {
	log.Infof("Delete subscriber %d", data.ChatId)

	var db []User
	var shouldWrite = false

	for index, v := range d {
		if v.ChatId == data.ChatId {
			shouldWrite = true
			db = removeElement(d, index)
		}
	}
	if shouldWrite {
		file, _ := json.MarshalIndent(db, "", "\t")
		_ = ioutil.WriteFile("database.json", file, 0644)
	}

}

func removeElement(s []User, i int) []User {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func listAll(path string) (photo map[int]string) {
	log.Infoln("List all photos...")

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

func ChoosePhotos(count int) (paths []string) {
	log.Infof("Choose %d photo(s)", count)
	photoMap := listAll(photos)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= count; i++ {
		index := rand.Intn(len(photoMap))
		paths = append(paths, photoMap[index])
		delete(photoMap, index)
	}

	return
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
