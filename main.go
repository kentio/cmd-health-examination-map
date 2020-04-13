package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kentio/cmd-health-examination/config"
	"github.com/kentio/cmd-health-examination/engine"
	"github.com/kentio/cmd-health-examination/model"
	"github.com/kentio/cmd-health-examination/parser"
)


func main() {
	db, err := model.InitDB()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	engine.Run(engine.Request{
		Url: config.BaseUrl + config.RootUrl,
		ParserFunc: parser.ParseCityList,
	})
	// 108 to json file -> data.json
	//var h []model.Hospital
	//r := make([]map[string]interface{},0)
	//
	//config.DataDB.Find(&h)
	//
	//for _,item := range h{
	//
	//	t := map[string]interface{}{
	//		"id" : item.ID,
	//		"name": item.Name,
	//		"city": item.City,
	//		"content": item.Content,
	//	}
	//	r = append(r, t)
	//}
	//fmt.Println(r)
	//
	//res,_ := json.Marshal(r)
	//f, _ := os.Create("data.json")
	//f.Write(res)
	//f.Close()


}
