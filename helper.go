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
	"strings"
	"time"
)

func readJSON() user {
	log.Debugf("Read json file...")
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
		currentJSON[id] = &userConfig{
			ChatId: id,
			Time:   []string{"18:11"},
		}

		saveJSON(currentJSON)
	}
}

func addMorePush(id int64, time string) (respond, message string) {
	log.Infof("Add more push for %d at %s", id, time)
	currentJSON := readJSON()

	currentPush := currentJSON[id].Time
	result := isContain(currentPush, time)
	if result {
		return "设置失败", "这个时间已经有了哦，小盆友你又调皮了呢😝"
	}

	currentPush = append(currentPush, time)
	currentJSON[id].Time = currentPush
	saveJSON(currentJSON)
	return "设置成功", "你现在的推送时间为 " + strings.Join(currentPush, " ")
}

func deleteOnePush(id int64, time string) {
	log.Infof("delete push entry for %d at %s", id, time)
	currentJSON := readJSON()

	currentPush := currentJSON[id].Time
	currentJSON[id].Time = removeElement(currentPush, time)
	saveJSON(currentJSON)

}
func removeElement(full []string, s string) (ret []string) {
	for _, v := range full {
		if v != s {
			ret = append(ret, v)
		}
	}
	return
}

func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func remove(id int64) {
	// delete all
	log.Infof("Delete subscriber %v", id)
	currentJSON := readJSON()
	delete(currentJSON, id)

	saveJSON(currentJSON)

}

func saveJSON(current user) {
	file, _ := json.MarshalIndent(current, "", "\t")
	log.Infof("Record json %v", current)
	err := ioutil.WriteFile("database.json", file, 0644)
	if err != nil {
		log.Errorf("Write json failed %v", err)
	}
}

func listAll(path string) (photo map[int]string) {
	log.Debugln("List all photos...")

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
	log.Debugf("Choose %d photo(s)", count)
	photoMap := listAll(photosPath)
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
	p := &tb.Photo{File: tb.FromDisk(chosen[0]), Caption: "怎么样，喜欢今日份的 Gakki 吗 😭😭😭"}
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
	if _, ok := config[uid]; ok {
		//存在
		return config[uid].Time
	} else {
		return []string{}
	}
}
