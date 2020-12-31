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

func readJSON() user {
	log.Infoln("Read json file...")
	jsonFile, _ := os.Open("database.json")
	decoder := json.NewDecoder(jsonFile)

	var config user
	err = decoder.Decode(&config)
	_ = jsonFile.Close()
	return config
}

func addInitSub(id int64) {
	log.Infof("Add subscriber %v", id)
	currentJSON := readJSON()

	// check if current user has already subscribed
	if _, ok := currentJSON[id]; !ok {
		// create user object
		currentJSON[id] = userConfig{
			ChatId: id,
			Time:   []string{"18:11"},
		}

		file, _ := json.MarshalIndent(currentJSON, "", "\t")
		log.Infoln("Record json %v", currentJSON)

		err := ioutil.WriteFile("database.json", file, 0644)
		if err != nil {
			log.Errorf("Write json failed %v", err)
		}

	}
}

func remove(id int64) {
	// delete all
	log.Infof("Delete subscriber %v", id)
	currentJSON := readJSON()
	delete(currentJSON, id)

	file, _ := json.MarshalIndent(currentJSON, "", "\t")
	log.Infoln("Record json %v", currentJSON)
	err := ioutil.WriteFile("database.json", file, 0644)
	if err != nil {
		log.Errorf("Write json failed %v", err)
	}

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
	photoMap := listAll(photosPath)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= count; i++ {
		index := rand.Intn(len(photoMap))
		paths = append(paths, photoMap[index])
		delete(photoMap, index)
	}

	log.Infof("Photo: %v", paths)
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

func timeSeries() (series []string) {
	var base int64 = 581983200
	for i := 0; i <= (22-7)*2; i++ {
		base += 60 * 30
		series = append(series, time.Unix(base, 0).Format("15:04"))
	}
	return series
}

func getPushTime(uid int64) []string {
	var config = readJSON()
	return config[uid].Time
}
