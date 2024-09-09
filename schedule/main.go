package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Total     int
	MatchList []GameResult
}

type GameResult struct {
	HomeID       string `json:"homeid"`
	H5PageID     string `json:"h5pageid"`
	StatusName   string `json:"statusname"`
	ItemCode     string `json:"itemcode"`
	SubItemName  string `json:"subitemname"`
	ID           string `json:"id"`
	VRTotalURL   string `json:"vrtotalurl"`
	EndDateCN    string `json:"enddatecn"`
	AlbumURL     string `json:"albumurl"`
	Title        string `json:"title"`
	VRLiveCode   string `json:"vrlivecode"`
	DeletedFlag  string `json:"deletedflag"`
	DocumentCode string `json:"documentcode"`
	TotalTitle   string `json:"totaltitle"`
	ImageURL     string `json:"imageurl"`
	VRLiveURL    string `json:"vrliveurl"`
	PageID       string `json:"pageid"`
	StartDateCN  string `json:"startdatecn"`
	VenueName    string `json:"venuename"`
	SubItemCode  string `json:"subitemcode"`
	AwayID       string `json:"awayid"`
	TotalGUID    string `json:"totalguid"`
	LockFlag     string `json:"lockflag"`
	Status       string `json:"status"`
	CombatFlag   string `json:"combatflag"`
	MVLiveCode   string `json:"mvlivecode"`
	HomeName     string `json:"homename"`
	LiveURL      string `json:"liveurl"`
	AwayName     string `json:"awayname"`
	Reserve3     string `json:"reserve3"`
	Venue        string `json:"venue"`
	AwayScore    string `json:"awayscore"`
	Reserve2     string `json:"reserve2"`
	HomeCode     string `json:"homecode"`
	Reserve1     string `json:"reserve1"`
	VRTotalCode  string `json:"vrtotalcode"`
	StartDate    string `json:"startdate"`
	HomeScore    string `json:"homescore"`
	LiveCode     string `json:"livecode"`
	ItemCodeName string `json:"itemcodename"`
	TotalURL     string `json:"totalurl"`
	MVLiveURL    string `json:"mvliveurl"`
	AwayCode     string `json:"awaycode"`
	AdCode       string `json:"adcode"`
	EndDate      string `json:"enddate"`
	Medal        string `json:"medal"`
}

type Detail struct {
	Name  string
	Time  string
	Sport string
	Venue string
}

func main() {

	dates := GetDates()

	for _, date := range dates {
		var ds []Detail
		url := "https://zy.api.cntv.cn/olympic/getOlyMatchList?itemcode=&startdate=2024" + date + "&venue=&medal=&t=jsonp&cb=OM&serviceId=2024aoyun&olyseason=2024S"
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		file, err := os.Create(date + ".json")
		if err != nil {
			log.Println(err)
		}

		buffer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		content := string(buffer)
		content = strings.TrimPrefix(content, "OM({\"data\":")
		content = strings.TrimSuffix(content, "});")

		data := new(Data)
		json.Unmarshal([]byte(content), data)

		for _, result := range data.MatchList {
			var d Detail
			d.Venue = result.VenueName
			d.Time = result.StartDateCN
			d.Sport = result.ItemCodeName
			if len(result.HomeScore) != 0 {
				d.Name = result.Title + " " + result.HomeName + "VS" + result.AwayName
			} else {
				d.Name = result.Title
			}
			//d.Name = result.Title + " " + result.HomeName + result.HomeScore + ":" + result.AwayScore + result.AwayName
			ds = append(ds, d)
		}

		ret, err := json.Marshal(ds)
		file.Write(ret)
		file.Close()
		resp.Body.Close()
	}

}

func GetDates() []string {
	var dates []string
	date := "07"

	for i := 24; i <= 31; i++ {
		dates = append(dates, date+strconv.Itoa(i))

	}
	date = "08"
	for i := 1; i <= 11; i++ {
		if i < 10 {
			dates = append(dates, date+"0"+strconv.Itoa(i))
		} else {
			dates = append(dates, date+strconv.Itoa(i))
		}
	}
	return dates
}
