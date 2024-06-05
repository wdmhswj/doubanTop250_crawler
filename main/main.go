package main

import (
	"crawler_doubanTop250/crawler"
	"crawler_doubanTop250/structs"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// list, err := crawler.GetTop250MovieData()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	list.SaveAsJson("../data/", "test0")
	// }

	var list1 structs.RankingList
	list1.LoadFromJson("../data/", "test0")
	fmt.Println("list1.name: ", list1.Name)
	fmt.Println("list1.url: ", list1.Url)

	var commentsForMovies []structs.MovieComments
	for i, movie := range list1.Movies {
		if i >= 50 {
			break
		}
		comments, err := crawler.GetComments(movie)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			commentsForMovies = append(commentsForMovies, *comments)
			// comments.SaveAsJson("../data/", "test_comments")
		}
	}
	// 将切片转换为 JSON 格式
	jsonData, err := json.MarshalIndent(commentsForMovies, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	filePath := "../data/comments_top50.json"
	// 将 JSON 数据写入文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("所有数据保存成功！！！")
}

func getMoviesAndSave() {
	list, err := crawler.GetTop250MovieData()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		list.SaveAsJson("../data/", "test0")
	}
}

func loadMovies() {
	var list1 structs.RankingList
	list1.LoadFromJson("../data/", "test0")
	fmt.Println("list1.name: ", list1.Name)
	fmt.Println("list1.url: ", list1.Url)
}
