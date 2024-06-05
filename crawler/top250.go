package crawler

import (
	"crawler_doubanTop250/structs"
	"fmt"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func getTop250MovieData() (*structs.RankingList, error) {
	var list structs.RankingList

	list.Name = "豆瓣电影Top250"
	list.Url = "https://movie.douban.com/top250"

	allowDomain := "movie.douban.com"

	c := colly.NewCollector(
		colly.AllowedDomains(allowDomain),
	)
	// // 设置自定义的 HTTP Transport 来处理压缩
	// c.WithTransport(&http.Transport{
	// 	DisableCompression: false,
	// })

	urlTemplate := "https://movie.douban.com/top250?start="

	// 在request之前
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())

		// 设置请求头
		r.Headers.Set("Connection", "keep-alive")
		// r.Headers.Set("Content-Encoding", "gzip")
		// r.Headers.Set("Server", "Lego Server")
		// r.Headers.Set("Cookie", "bid=HRkyp9pX-D8; ll=\"118164\"; douban-fav-remind=1; _pk_id.100001.4cf6=6ca3cc9d6dabf1c7.1707124210.; _vwo_uuid_v2=DF952C2B33E4731C6F27611DAE0EC925C|8b977aed23f4803548ba92e6658832e7; __utmz=223695111.1707130612.2.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; viewed=\"35128111_25868289_10433737_30153583_26323176_35750846_1627194_35751619_35751623_34898994\"; __utmz=30149280.1717152507.14.8.utmcsr=bing|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); ap_v=0,6.0; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1717584400%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_ses.100001.4cf6=1; __utma=30149280.886874396.1704267674.1717152507.1717585680.15; __utmb=30149280.0.10.1717585680; __utmc=30149280; __utma=223695111.729703034.1707124210.1707130612.1717585680.3; __utmb=223695111.0.10.1717585680; __utmc=223695111")
		r.Headers.Set("Host", "movie.douban.com")
		r.Headers.Set("Referer", "https://movie.douban.com/top250?start=0")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
		r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"124\", \"Microsoft Edge\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		// r.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	})

	// 收到响应Response时
	c.OnResponse(func(r *colly.Response) {
		// 打印响应的状态码
		fmt.Println("响应状态码:", r.StatusCode)

		// 打印响应头
		fmt.Println("响应头:", r.Headers)

		// 打印响应体（HTML 或者其他数据）
		// fmt.Println("响应体:", string(r.Body))
		fmt.Println("响应体: 。。。")
	})

	c.OnHTML("div.article ol.grid_view li", func(e *colly.HTMLElement) {
		var movie structs.Movie
		var err error

		// 排名
		index := e.DOM.Find("div.item div.pic em").Text()
		movie.TopIndex, err = strconv.Atoi(index)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 名称
		movie.MovieName = e.DOM.Find("div.item div.info div.hd a span.title").Text()

		// 总评分
		rating_num := e.DOM.Find("div.item div.info div.bd div.star span.rating_num").Text()
		float64Num, err := strconv.ParseFloat(rating_num, 32)
		if err != nil {
			fmt.Println("Error:", err)
		}
		movie.TotalScore = float32(float64Num)

		// 详情页url
		movie.DetailPageUrl, _ = e.DOM.Find("div.item div.info div.hd a").Attr("href")

		fmt.Println("完成1个实体的爬取：")
		fmt.Println("\t名称：", movie.MovieName)
		fmt.Println("\t评分：", movie.TotalScore)
		fmt.Println("\t排名：", movie.TopIndex)
		fmt.Println("\t详情页url", movie.DetailPageUrl)

		list.Movies = append(list.Movies, movie)
	})

	var pageNumber = 10
	for i := 0; i < pageNumber; i++ {
		url := urlTemplate + strconv.Itoa(25*i)
		c.Visit(url)
	}

	return &list, nil
}

var GetTop250MovieData = getTop250MovieData
