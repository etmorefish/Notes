package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "123"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "demo"
)

var DB *sql.DB

type MovieData struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Picture  string `json:"picture"`
	Actor    string `json:"actor"`
	Year     string `json:"year"`
	Score    string `json:"score"`
	Quote    string `json:"quote"`
}

func Spider(page string) {
	//1 发送请求
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250?start="+page, nil)
	if err != nil {
		fmt.Println("req accept err:", err)
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp accept err:", err)
	}
	defer resp.Body.Close()
	//2 解析网页
	docDetails, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("解析失败", err)
	}

	//3 获取节点信息
	// #content > div > div.article > ol > li:nth-child(1)
	// #content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	docDetails.Find("#content > div > div.article > ol > li > div").
		Each(func(i int, s *goquery.Selection) {
			var movieData MovieData

			title := s.Find("div.info > div.hd > a > span:nth-child(1)").Text()
			img := s.Find("div.pic > a > img")
			imgTmp, ok := img.Attr("src")
			info := strings.Trim(s.Find("div.info > div.bd > p:nth-child(1)").Text(), " ")
			director, actor, year := InfoSpite(info)
			score := strings.Trim(s.Find("div.info > div.bd > div > span.rating_num").Text(), " ")
			score = strings.Trim(score, "\n")
			quote := strings.Trim(s.Find("div.info > div.bd > p.quote > span").Text(), " ")

			if ok {

				movieData.Title = title
				movieData.Director = director
				movieData.Picture = imgTmp
				movieData.Actor = actor
				movieData.Year = year
				movieData.Score = score
				movieData.Quote = quote

				if InsertSql(movieData) {

				} else {
					fmt.Println("insert err")
				}
				// fmt.Println(movieData)

			}
		})

	fmt.Println("insert success")
	return
	//4 保存信息
}

func InfoSpite(info string) (director, actor, year string) {
	directorRe, _ := regexp.Compile(`导演:(.*)主演:`)
	director = string(directorRe.Find([]byte(info)))
	actorRe, _ := regexp.Compile(`主演:(.*)`)
	actor = string(actorRe.Find([]byte(info)))
	yearRe, _ := regexp.Compile(`(\d+)`)
	year = string(yearRe.Find([]byte(info)))
	return
}

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertSql(movieData MovieData) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail", err)
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO movie_data (`Title`,`Director`,`Picture`,`Actor`,`Year`,`Score`,`Quote`) VALUES (?, ?, ?,?,?,?,?)")
	if err != nil {
		fmt.Println("Prepare fail", err)
	}
	_, err = stmt.Exec(movieData.Title, movieData.Director, movieData.Picture, movieData.Actor, movieData.Year, movieData.Score, movieData.Quote)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	_ = tx.Commit()
	return true
}

func main() {
	InitDB()
	for i := 0; i < 10; i++ {
		fmt.Printf("正在爬取第%d页", i)
		Spider(strconv.Itoa(i * 25))
	}
}
