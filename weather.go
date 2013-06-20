package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	showHelp = flag.Bool("h", false, "show help")
	cityCode = map[string]string{
		"nj": "101190101",
		"sz": "101190401",
	}
)

func main() {
	flag.Parse()

	//must have at least one word
	if flag.NArg() == 0 || *showHelp {
		usage()
		os.Exit(1)
	}

	getWeather(getCityCode(flag.Args()))
}

func usage() {
	fmt.Printf("%s city1 [city2...cityN]\n", os.Args[0])
	flag.PrintDefaults()
}

func getCityCode(citys []string) []string {
	codes := make([]string, 0, len(citys))
	for _, city := range citys {
		if code, ok := cityCode[city]; ok {
			codes = append(codes, code)
		} else {
			log.Printf("city %q is not supported\n", city)
		}
	}
	return codes
}

type WeatherInfo struct {
	City     string `json:"city"`
	Date     string `json:"date_y"`
	Temp1    string `json:"temp1"`
	Temp2    string `json:"temp2"`
	Temp3    string `json:"temp3"`
	Temp4    string `json:"temp4"`
	Temp5    string `json:"temp5"`
	Temp6    string `json:"temp6"`
	Weather1 string `json:"weather1"`
	Weather2 string `json:"weather2"`
	Weather3 string `json:"weather3"`
	Weather4 string `json:"weather4"`
	Weather5 string `json:"weather5"`
	Weather6 string `json:"weather6"`
}

type Weather struct {
	Info WeatherInfo `json:"weatherinfo"`
}

func (w *Weather) String() (result string) {
	wi := w.Info
	result += fmt.Sprintf("%s, %s, %s, %s\n",
		wi.City, wi.Date, wi.Weather1, wi.Temp1)
	result += fmt.Sprintf("futher 5 weathers:\n%s, %s\n%s, %s\n%s, %s\n%s, %s\n%s, %s\n",
		wi.Weather2, wi.Temp2, wi.Weather3, wi.Temp3, wi.Weather4, wi.Temp4, wi.Weather5, wi.Temp5, wi.Weather6, wi.Temp6)
	return
}

func getWeather(codes []string) {
	const urlPrefix = "http://m.weather.com.cn/data/"
	for _, code := range codes {
		resp, err := http.Get(urlPrefix + code + ".html")
		if err != nil {
			log.Println(err)
			continue
		}
		resultJson, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println(err)
			continue
		}

		weather := new(Weather)
		if err := json.Unmarshal(resultJson, weather); err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(weather)
	}
}
