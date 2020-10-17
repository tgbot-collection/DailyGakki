// DailyGakki - helper
// 2020-10-17 16:37
// Benny <benny.think@gmail.com>

package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func readJSON() []User {

	jsonFile, _ := os.Open("database.json")
	decoder := json.NewDecoder(jsonFile)

	var db []User
	err = decoder.Decode(&db)
	_ = jsonFile.Close()
	return db

}

func add(d []User, data User) {
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
	photoMap := listAll(photos)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= count; i++ {
		index := rand.Intn(len(photoMap))
		paths = append(paths, photoMap[index])
		delete(photoMap, index)
	}

	return
}
