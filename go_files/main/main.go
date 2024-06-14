package main

import (
	"crawler_doubanTop250/crawler"
	"crawler_doubanTop250/structs"
	"crawler_doubanTop250/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	csvfile := "../data/merged_comments_2col_v3.csv"
	utils.CsvStatistics(csvfile)

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

	var commentsForMovies []structs.MovieComments
	for i, movie := range list1.Movies {
		if i < 149 {
			continue
		}
		fmt.Printf("%d. %s(%s)\n", i, movie.MovieName, movie.DetailPageUrl)
		comments, err := crawler.GetComments(movie)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			commentsForMovies = append(commentsForMovies, *comments)
			// comments.SaveAsJson("../data/", "test_comments")
		}
		// break
	}
	// 将切片转换为 JSON 格式
	jsonData, err := json.MarshalIndent(commentsForMovies, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	filePath := "../data/comments_top150-250.json"
	// 将 JSON 数据写入文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("所有数据保存成功！！！")
}

// 读取 JSON 文件并解析为结构体
func readJSONFile(filePath string) ([]structs.MovieComments, error) {
	var movies []structs.MovieComments
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

// 将合并后的切片保存为 CSV 文件
func saveAsCSV(filePath string, movies []structs.MovieComments) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入CSV文件的头
	err = writer.Write([]string{"UserComment", "UserRating"})
	if err != nil {
		return err
	}

	// 写入电影评论数据
	for _, movie := range movies {
		for _, comment := range movie.Comments {
			if comment.UserRating == 0 {
				continue
			}
			// commentText := strings.Replace(comment.UserComment, "\n", "", -1)
			commentText := strings.ReplaceAll(comment.UserComment, "\n", "")
			commentText = strings.ReplaceAll(commentText, "\r", "")
			commentText = strings.ReplaceAll(commentText, "\t", "")
			// commentText = strings.ReplaceAll(commentText, " ", "")
			record := []string{
				commentText,
				strconv.Itoa(comment.UserRating),
			}
			err = writer.Write(record)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func saveAsCSV_2class(filePath string, movies []structs.MovieComments) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入CSV文件的头
	err = writer.Write([]string{"MovieName", "MoiveIndex", "UserRating", "UserComment"})
	if err != nil {
		return err
	}

	// 写入电影评论数据
	for _, movie := range movies {
		for _, comment := range movie.Comments {
			if comment.UserRating == 3 {
				continue
			} else if comment.UserRating < 3 {
				record := []string{
					movie.MovieName,
					strconv.Itoa(movie.MoiveIndex),
					"0",
					comment.UserComment,
				}
				err = writer.Write(record)
				if err != nil {
					return err
				}
			} else {
				record := []string{
					movie.MovieName,
					strconv.Itoa(movie.MoiveIndex),
					"1",
					comment.UserComment,
				}
				err = writer.Write(record)
				if err != nil {
					return err
				}
			}
			// record := []string{
			// 	movie.MovieName,
			// 	strconv.Itoa(movie.MoiveIndex),
			// 	strconv.Itoa(comment.UserRating),
			// 	comment.UserComment,
			// }
			// err = writer.Write(record)
			// if err != nil {
			// 	return err
			// }
		}
	}

	return nil
}

func saveAsCSV_3class(filePath string, movies []structs.MovieComments) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入CSV文件的头
	err = writer.Write([]string{"MovieName", "MoiveIndex", "UserRating", "UserComment"})
	if err != nil {
		return err
	}

	// 写入电影评论数据
	for _, movie := range movies {
		for _, comment := range movie.Comments {
			if comment.UserRating == 3 {
				record := []string{
					movie.MovieName,
					strconv.Itoa(movie.MoiveIndex),
					"1",
					comment.UserComment,
				}
				err = writer.Write(record)
				if err != nil {
					return err
				}
			} else if comment.UserRating < 3 {
				record := []string{
					movie.MovieName,
					strconv.Itoa(movie.MoiveIndex),
					"0",
					comment.UserComment,
				}
				err = writer.Write(record)
				if err != nil {
					return err
				}
			} else {
				record := []string{
					movie.MovieName,
					strconv.Itoa(movie.MoiveIndex),
					"2",
					comment.UserComment,
				}
				err = writer.Write(record)
				if err != nil {
					return err
				}
			}
			// record := []string{
			// 	movie.MovieName,
			// 	strconv.Itoa(movie.MoiveIndex),
			// 	strconv.Itoa(comment.UserRating),
			// 	comment.UserComment,
			// }
			// err = writer.Write(record)
			// if err != nil {
			// 	return err
			// }
		}
	}

	return nil
}
func saveAsJson(movies []structs.MovieComments) {
	// 序列化合并后的数组为 JSON 格式
	mergedJSONData, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Println("序列化 JSON 失败:", err)
		return
	}

	// 将合并后的 JSON 数据写入一个新的文件
	err = os.WriteFile("../data/merged_comments.json", mergedJSONData, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败:", err)
		return
	}

	fmt.Println("成功将两个 JSON 文件的数据合并并写入新的 JSON 文件！")
}

// 验证CSV文件格式
func validateCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// 读取CSV文件的头
	_, err = reader.Read()
	if err != nil {
		return err
	}

	// 读取和验证CSV文件的数据
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		if len(record) != 2 {
			return fmt.Errorf("invalid record length: %d", len(record))
		}

		// 验证UserComment为字符串
		// userComment := record[0]
		// if _, err := strconv.Atoi(userComment); err == nil {
		// 	return fmt.Errorf("userComment should be a string, got a number: %s", userComment)
		// }

		// 验证UserRating为整数
		_, err = strconv.Atoi(record[1])
		if err != nil {
			return fmt.Errorf("userRating should be an integer, got: %s", record[1])
		}

		// 进一步的验证可以在这里添加
		// fmt.Printf("Validated comment: %s, rating: %d\n", userComment, userRating)
	}

	return nil
}

func merger2json() {
	// 读取第一个 JSON 文件
	movies1, err := readJSONFile("../data/comments_top1-149.json")
	if err != nil {
		fmt.Println("读取 JSON 文件失败:", err)
		return
	}

	// 读取第二个 JSON 文件
	movies2, err := readJSONFile("../data/comments_top150-250.json")
	if err != nil {
		fmt.Println("读取 JSON 文件失败:", err)
		return
	}

	// fmt.Println(movies1[0])
	// 合并两个数组
	movies := append(movies1, movies2...)

	// 定义输出文件路径
	outputCSVFilePath := "../data/merged_comments_2col_v3.csv"

	// 将合并后的切片保存为 CSV 文件
	err = saveAsCSV(outputCSVFilePath, movies)
	if err != nil {
		fmt.Println("写入 CSV 文件失败:", err)
		return
	}

	err = validateCSV(outputCSVFilePath)
	if err != nil {
		fmt.Println("Error validating CSV:", err)
	} else {
		fmt.Println("CSV file validated successfully.")
	}
}
