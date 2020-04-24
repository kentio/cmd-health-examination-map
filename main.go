package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kentio/cmd-health-examination-map/config"
	"github.com/kentio/cmd-health-examination-map/engine"
	"github.com/kentio/cmd-health-examination-map/model"
	"github.com/kentio/cmd-health-examination-map/parser"
	"github.com/kentio/cmd-health-examination-map/scheduler"
	"os"
	"strings"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        config.BaseUrl + config.RootUrl,
		ParserFunc: parser.ParseCityList,
	})

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        config.BaseUrl + config.RootUrl,
	//	ParserFunc: parser.ParseCityList,
	//})

	// 108 to json file -> data.json
	//dataToJson()

}

func dataToJson() {
	var h []model.Hospital
	r := make([]map[string]interface{}, 0)

	config.DataDB.Find(&h)

	for _, item := range h {
		var addr, tel string

		content := strings.Split(item.Content, "\n")
		fmt.Println(content)
		for _, item := range content {
			if ok := strings.Contains(item, "门店地址"); ok {
				addr = item
			}
			if ok := strings.Contains(item, "联系电话"); ok {
				tel = item
			}
		}

		t := map[string]interface{}{
			"id":      item.ID,
			"name":    item.Name,
			"city":    item.City,
			"content": item.Content,
			"addr":    addr,
			"tel":     tel,
		}
		r = append(r, t)
	}
	//fmt.Println(r)
	// save to file
	res, _ := json.Marshal(r)
	f, _ := os.Create("data.json")
	f.Write(res)
	f.Close()
}
