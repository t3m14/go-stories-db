package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)
var DB *gorm.DB

type Story struct {
  ID         int "gorm:'id'"
  Hidden     bool   "gorm:'hidden'"
  BusinessId int "gorm:'businessId'"
  CreatedAt  string "gorm:'createdAt'"
  Story      string "gorm:'story'"
  Views      string "gorm:'views'"
  Likes      string    "gorm:'likes'"
}
func main() {
	fmt.Println("Initializing...")
	InitDb()
}
// Инициализируем базу данных, можно заменить на любую другую на строке ниже
func InitDb(){
	db, err := gorm.Open("mysql",
		"root:root@/stories")
	
	if err != nil {
	fmt.Println(err)
	return
	} else {
	fmt.Println("Connection success...")
	}
	defer db.Close()
	DB = db
}

// Функция для добавления новой истории
func add_story(story_y *Story)(id int, hidden bool, businessId int, createdAt string, story string, views string, likes string) {
  story_y = &Story{ID: id, Hidden : hidden, BusinessId:businessId, CreatedAt:createdAt, Story:story, Views:views, Likes:likes,}
  DB.Create(story)
  return
}
//Функция удаления стори по айди
func delete(id int){
	story := &Story{ID:id}
	DB.Delete(story)
}
//Функция обновления записи из базы данных по айдиp, ключу и значению.
func update_story(id int, key string, value string){
	story := &Story{ID:id}
	DB.Model(story).Update(key, value)
}
