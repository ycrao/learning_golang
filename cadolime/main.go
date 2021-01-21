package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

func main() {
	ld := LionDeltaTime()
	dd := DogDeltaTime()
	cd := CatDeltaTime()
	SystemTime(ld, dd, cd)
}

func SystemTime(ld, dd, cd int64) {
	for {
		sTs := time.Now().UnixNano()/1e6

		lTs := sTs + ld
		dTs := sTs + dd
		cTs := sTs + cd

		stm := time.Unix(0, sTs*int64(time.Millisecond))

		ltm := time.Unix(0, lTs*int64(time.Millisecond))
		dtm := time.Unix(0, dTs*int64(time.Millisecond))
		ctm := time.Unix(0, cTs*int64(time.Millisecond))

		fmt.Printf("\rsytemTime:%s", stm.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("\rlionTime:%s", ltm.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("\rdogTime:%s", dtm.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("\rcatTime:%s", ctm.Format("2006-01-02 15:04:05.000"))
		time.Sleep(100 * time.Millisecond)
	}
}

func CatDeltaTime() int64 {
	cTs, _ := fetchTime("cat")
	sTs := time.Now().UnixNano() / 1e6  // system time
	delta := sTs - cTs
	fmt.Println("cat delta:", delta)
	return delta
}

func DogDeltaTime() int64 {
	cTs, _ := fetchTime("dog")
	sTs := time.Now().UnixNano() / 1e6  // system time
	delta := sTs - cTs
	fmt.Println("dog delta:", delta)
	return delta
}

func LionDeltaTime() int64 {
	uTs, _ := fetchTime("lion")
	// fmt.Println("start", uTs)
	iTs := uTs
	count := 0
	var c1, c2, dCount int
	var lastTs, sTs, delta int64
	Loop:
		for {
			lastTs = uTs
			uTs, _ = fetchTime("lion")
			count ++
			if (lastTs != uTs)  && (uTs == iTs + 1) {
				c1 = count
			} else if (lastTs != uTs) && (uTs == iTs + 2) {
				c2 = count
				sTs = time.Now().UnixNano() / 1e6  // system time
				dCount = c2 - c1
				requestSpendTime := 1000/dCount
				msUts := uTs*1000 + int64(requestSpendTime)  // added network spend time per request
				delta = sTs - msUts
				// fmt.Println(sTs, msUts)
				break Loop
			}
			time.Sleep(40*time.Millisecond)
		}
	// fmt.Println(iTs, uTs, sTs, c1, c2, delta)
	fmt.Println("lion delta:", delta)
	return delta
}

const (
	CAT_URL = "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"  // ms timestamp
	DOG_URL = "https://a.jd.com//ajax/queryServerData.html"  // second timestamp
	LION_URL = "http://quan.suning.com/getSysTime.do"  // millisecond timestamp
)

/*
```
#cat's time
{
	"api": "mtop.common.getTimestamp",
	"v": "*",
	"ret": [
		"SUCCESS::接口调用成功"
	],
	"data": {
		"t": "1610778106791"
	}
}

#dog's time
{
	"serverTime": 1610782350944
}

#lion's time
{
	"sysTime2": "2021-01-16 15:32:39",
	"sysTime1": "20210116153239"
}
```
 */
func fetchTime(animal string) (int64, string) {
	var url string
	switch animal {
	case "dog":
		url = DOG_URL
		resp, _ := resty.New().R().Get(url)
		ts := gjson.Get(string(resp.Body()), "serverTime").Int()
		datetimeStr := time.Unix(ts, 0).Format("2006-01-02 15:04:05")
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	case "lion":
		url = LION_URL
		resp, _ := resty.New().R().Get(url)
		datetimeStr := gjson.Get(string(resp.Body()), "sysTime2").String()
		oTime, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.Local)
		ts := oTime.Unix()
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	default:
		url = CAT_URL
		resp, _ := resty.New().R().Get(url)
		tsStr := gjson.Get(string(resp.Body()), "data.t").String()
		ts, _ := strconv.ParseInt(tsStr, 10, 64)
		datetimeStr := time.Unix(ts, 0).Format("2006-01-02 15:04:05")
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	}
}