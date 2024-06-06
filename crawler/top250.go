package crawler

import (
	"crawler_doubanTop250/structs"
	"fmt"
	"strconv"
	"strings"

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

func getComments(movie structs.Movie) (*structs.MovieComments, error) {
	var comments structs.MovieComments
	comments.MovieName = movie.MovieName
	comments.MoiveIndex = movie.TopIndex

	// url模板
	movieUrl := movie.DetailPageUrl
	// 热门（new_score），好评（percent_type=h）
	// https://movie.douban.com/subject/26615208/comments?percent_type=h&start=*&limit=20&status=P&sort=new_score
	// https://movie.douban.com/subject/1292052/
	urlTemplate_hot_h := movieUrl + "comments?percent_type=h&start=*#*&limit=20&status=P&sort=new_score"
	// 热门，一般（percent_type=m）
	// https://movie.douban.com/subject/26615208/comments?percent_type=m&limit=20&status=P&sort=new_score
	urlTemplate_hot_m := movieUrl + "comments?percent_type=m&start=*#*&limit=20&status=P&sort=new_score"
	// 热门，一般（percent_type=l）
	// https://movie.douban.com/subject/26615208/comments?percent_type=l&limit=20&status=P&sort=new_score
	urlTemplate_hot_l := movieUrl + "comments?percent_type=l&start=*#*&limit=20&status=P&sort=new_score"

	// 最新（sort=time）
	// https://movie.douban.com/subject/26615208/comments?start=*&limit=20&status=P&sort=time
	urlTemplate_new := movieUrl + "comments?start=*#*&limit=20&status=P&sort=time"

	allowDomain := "movie.douban.com"
	c := colly.NewCollector(
		colly.AllowedDomains(allowDomain),
	)

	// 在request之前
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting:", r.URL.String())

		// 设置请求头
		r.Headers.Set("Connection", "keep-alive")
		// r.Headers.Set("Content-Encoding", "gzip")
		// r.Headers.Set("Server", "Lego Server")
		r.Headers.Set("Cookie", "bid=HRkyp9pX-D8; ll=\"118164\"; douban-fav-remind=1; _pk_id.100001.4cf6=6ca3cc9d6dabf1c7.1707124210.; _vwo_uuid_v2=DF952C2B33E4731C6F27611DAE0EC925C|8b977aed23f4803548ba92e6658832e7; __utmz=223695111.1707130612.2.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; viewed=\"35128111_25868289_10433737_30153583_26323176_35750846_1627194_35751619_35751623_34898994\"; __utmz=30149280.1717152507.14.8.utmcsr=bing|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmc=30149280; __utmc=223695111; ap_v=0,6.0; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1717597775%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_ses.100001.4cf6=1; __utma=30149280.886874396.1704267674.1717590470.1717597776.17; __utmt=1; __utma=223695111.729703034.1707124210.1717590470.1717597809.5; __utmb=223695111.0.10.1717597809; __utmb=30149280.7.10.1717597776; dbcl2=\"280989070:bXH6OaO894k\"; ck=Xh55; push_noty_num=0; push_doumail_num=0")
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
		// fmt.Println("响应状态码:", r.StatusCode)

		// 打印响应头
		// fmt.Println("响应头:", r.Headers)

		// 打印响应体（HTML 或者其他数据）
		// fmt.Println("响应体:", string(r.Body))
		// fmt.Println("响应体: 。。。")
	})

	c.OnHTML("div.comment", func(e *colly.HTMLElement) {
		var comment structs.UserShortComment

		// 获取评分
		ratingClass := e.ChildAttr("span[class^='allstar']", "class")
		if ratingClass != "" {
			// fmt.Println("test1")
			ratingStr := strings.TrimPrefix(ratingClass, "allstar")
			// fmt.Println(ratingStr)
			ratingNum, err := strconv.Atoi(ratingStr[:2])
			if err == nil {
				comment.UserRating = ratingNum / 10
				// fmt.Println("test2")
			}
		}
		// 获取短评内容
		comment.UserComment = e.ChildText("span.short")

		// fmt.Println("获取到1个评论：")
		// fmt.Println("\t评分：", comment.UserRating)
		// fmt.Println("\t短评：", comment.UserComment)

		comments.Comments = append(comments.Comments, comment)
	})

	// 热门，好评
	for i := 0; i < 10; i++ {
		url := strings.Replace(urlTemplate_hot_h, "*#*", fmt.Sprintf("%d", i*20), 1)
		// fmt.Println("url to visit:", url)
		c.Visit(url)
	}
	// 热门，一般
	for i := 0; i < 10; i++ {
		url := strings.Replace(urlTemplate_hot_m, "*#*", fmt.Sprintf("%d", i*20), 1)
		// fmt.Println("url to visit:", url)
		c.Visit(url)
	}
	// 热门，差评
	for i := 0; i < 10; i++ {
		url := strings.Replace(urlTemplate_hot_l, "*#*", fmt.Sprintf("%d", i*20), 1)
		// fmt.Println("url to visit:", url)
		c.Visit(url)
	}
	// 最新
	for i := 0; i < 30; i++ {
		url := strings.Replace(urlTemplate_new, "*#*", fmt.Sprintf("%d", i*20), 1)
		// fmt.Println("url to visit:", url)
		c.Visit(url)
	}

	return &comments, nil
}

var GetComments = getComments
