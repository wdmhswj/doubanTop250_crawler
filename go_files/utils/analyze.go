package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CsvStatistics(filepath string) {
	// 打开CSV文件
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个CSV读取器
	reader := csv.NewReader(file)

	// 读取CSV文件的所有内容
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var sum, count, maxCommentLen, minCommentLen int
	minCommentLen = int(^uint(0) >> 1) // 设置为最大整数
	// 创建一个映射来统计每个评分的出现次数
	ratingCounts := make(map[int]int)

	// 遍历读取的记录
	for _, record := range records {
		// 假设CSV文件的第一列是UserComment，第二列是UserRating
		userComment := record[0]
		userRatingStr := record[1]

		// 将UserRating从字符串转换为整数
		userRating, err := strconv.Atoi(userRatingStr)
		if err != nil {
			log.Printf("Invalid rating %s: %v", userRatingStr, err)
			continue
		}

		// 检查UserRating是否在1到5的范围内
		if userRating >= 1 && userRating <= 5 {
			// 统计评分
			ratingCounts[userRating]++
			sum += userRating
			count++
		} else {
			log.Printf("Rating out of range: %d", userRating)
		}
		// 统计评论长度
		commentLen := len(userComment)
		if commentLen > maxCommentLen {
			maxCommentLen = commentLen
		}
		if commentLen < minCommentLen {
			minCommentLen = commentLen
			fmt.Println(userComment)
		}
		// fmt.Printf("User Comment: %s, User Rating: %d\n", userComment, userRating)
	}

	// 计算平均值
	var mean float64
	if count != 0 {
		mean = float64(sum) / float64(count)
	}

	// 打印评分统计结果
	fmt.Println("Rating Statistics:")
	for rating := 1; rating <= 5; rating++ {
		fmt.Printf("Rating %d: %d\n", rating, ratingCounts[rating])
	}
	fmt.Printf("Average Rating: %.2f\n", mean)
	fmt.Printf("Total Ratings: %d\n", count)

	// 打印评论长度统计结果
	fmt.Printf("Longest Comment Length: %d\n", maxCommentLen)
	fmt.Printf("Shortest Comment Length: %d\n", minCommentLen)

}
