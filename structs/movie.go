package structs

import (
	"crawler_doubanTop250/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Movie struct {
	MovieName  string  // 电影名称
	TotalScore float32 // 电影总评分
	// MovieType     string			// 电影种类
	DetailPageUrl string // 电影详情页
	TopIndex      int    // 排行
}
type UserShortComment struct {
	MovieName   string
	UserRating  int
	UserComment string
}

type RankingList struct {
	Name   string // 榜单名称
	Url    string // 榜单首页url
	Movies []Movie
}

func (list *RankingList) SaveAsJson(saveDir string, filename string) {
	// 若不存在data目录则创建
	ok, err := utils.FileDirExist(saveDir)
	if !ok {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = os.Mkdir(saveDir, 0750)
		if err != nil {
			fmt.Println("failed to mkdir:", err.Error())
			return
		}
	}
	filename += ".json"
	filename = filepath.Join(saveDir, filename)
	// filename = dir + filename + ".json"
	// fmt.Println(filename)
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("相同文件名称的JSON文件已存在！")
	} else {

		// 将结构体实例序列化为 JSON 格式
		jsonData, err := json.MarshalIndent(list, "", "    ")
		if err != nil {
			fmt.Println("序列化 JSON 失败:", err)
			return
		}

		// 将 JSON 数据写入本地文件
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
			fmt.Println("写入 JSON 文件失败:", err)
			return
		}
		fmt.Println("JSON 文件保存成功！")
	}
}
