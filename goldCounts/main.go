package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Rank struct {
	Bronze      string
	Rank        string
	Count       string
	Silver      string
	Countryname string
	Gold        string
	Countryid   string
}

type Data struct {
	Total      int
	MedalsList []Rank
}

func main() {
	data := new(Data)

	goldRankUrl := "https://zy.api.cntv.cn/Olympic/getOlyMedals?serviceId=2024aoyun&olyseason=2024S&itemcode=GEN-------------------------------&t=jsonp&cb=jpb"
	resp, err := http.Get(goldRankUrl)
	if err != nil {
		log.Println(err)
	}
	//ReadData(resp)
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	content := string(buffer)
	c1 := strings.TrimPrefix(content, "jpb({\"data\":")
	c2 := strings.TrimSuffix(c1, "});")
	file, err := os.Create("gold_rank.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	json.Unmarshal([]byte(c2), &data)

	ret, err := json.Marshal(data)
	//log.Println(data)
	file.Write(ret)

}

func ReadData(resp *http.Response) {
	file, err := os.Create("gold_rank.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Println(err)
	}
	doc.Find("tbody").Each(func(i int, s1 *goquery.Selection) {
		v1, e1 := s1.Attr("id")
		if e1 && v1 == "medal_list1" {

			s1.Find("td").Each(func(j int, s2 *goquery.Selection) {
				file.WriteString(s2.Text())
			})
		}
	})
}
