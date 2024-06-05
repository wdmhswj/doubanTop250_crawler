package structs

import (
	"crawler_doubanTop250/utils"
	"encoding/json"
	"fmt"
	"io"
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
	// MovieName   string
	UserRating int
	// IsGoodReview int
	UserComment string
}

type RankingList struct {
	Name   string // 榜单名称
	Url    string // 榜单首页url
	Movies []Movie
}

// 电影短评及评分的容器
type MovieComments struct {
	MovieName  string
	MoiveIndex int
	Comments   []UserShortComment
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

func (list *RankingList) LoadFromJson(saveDir string, filename string) {
	// 若不存在data目录则创建
	ok, _ := utils.FileDirExist(saveDir)
	if !ok {
		fmt.Println("Dir does not exist!")
		return
	}
	filename += ".json"
	filename = filepath.Join(saveDir, filename)
	// filename = dir + filename + ".json"
	// fmt.Println(filename)
	if _, err := os.Stat(filename); err != nil {
		fmt.Println(filename + "对应的JSON文件不存在！")
		return
	} else {
		// 打开 JSON 文件
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("打开 JSON 文件失败:", err)
			return
		}
		defer file.Close()

		// 读取 JSON 数据
		jsonData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("读取 JSON 数据失败:", err)
			return
		}

		err = json.Unmarshal(jsonData, list)
		if err != nil {
			fmt.Println("反序列化 JSON 失败:", err)
			return
		}

		fmt.Println("反序列化成功")
		return

	}
}

func (comments *MovieComments) SaveAsJson(saveDir string, filename string) {
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

		// 打开文件以追加模式
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return
		}
		defer file.Close()

		// 将结构体实例序列化为 JSON 格式
		jsonData, err := json.MarshalIndent(comments, "", "    ")
		if err != nil {
			fmt.Println("序列化 JSON 失败:", err)
			return
		}

		// 写入 JSON 数据到文件末尾
		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("写入 JSON 数据失败:", err)
			return
		}
		fmt.Println("JSON 数据追加到文件末尾成功！")

	} else {

		// 将结构体实例序列化为 JSON 格式
		jsonData, err := json.MarshalIndent(comments, "", "    ")
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

// TODO
func (comments *MovieComments) LoadFromJson(saveDir string, filename string) {
	// 若不存在data目录则创建
	ok, _ := utils.FileDirExist(saveDir)
	if !ok {
		fmt.Println("Dir does not exist!")
		return
	}
	filename += ".json"
	filename = filepath.Join(saveDir, filename)
	// filename = dir + filename + ".json"
	// fmt.Println(filename)
	if _, err := os.Stat(filename); err != nil {
		fmt.Println(filename + "对应的JSON文件不存在！")
		return
	} else {
		// 打开 JSON 文件
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("打开 JSON 文件失败:", err)
			return
		}
		defer file.Close()

		// 读取 JSON 数据
		jsonData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("读取 JSON 数据失败:", err)
			return
		}

		err = json.Unmarshal(jsonData, comments)
		if err != nil {
			fmt.Println("反序列化 JSON 失败:", err)
			return
		}

		fmt.Println("反序列化成功")
		return

	}
}
