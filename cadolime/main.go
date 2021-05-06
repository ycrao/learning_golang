package main

import (
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

var (
	p = flag.String("p", "sys", "platform: sys/lion/dog/dog")
	d = flag.Int64("d", 10, "request delay: ping platform website using avg ms time")
)

const (
	CAT_URL = "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"  // millisecond timestamp
	DOG_URL = "https://a.jd.com//ajax/queryServerData.html"  // millisecond timestamp
	LION_SECOND_URL = "http://quan.suning.com/getSysTime.do"  // second timestamp [please control frequency]
	LION_URL = "https://f.m.suning.com/api/ct.do"  // millisecond timestamp
)

func main() {
	flag.Parse()
	ld := LionDeltaTime()
	dd := DogDeltaTime()
	cd := CatDeltaTime()
	// fmt.Println(ld, dd, cd)
	SystemTime(ld, dd, cd)
}

func SystemTime(ld, dd, cd int64) {
	platform := *p
	delay := *d
	for {
		sTs := time.Now().UnixNano()/1e6
		lTs := sTs - ld + delay
		dTs := sTs - dd + delay
		cTs := sTs - cd + delay
		// stm := time.Unix(0, sTs*int64(time.Millisecond))
		ltm := time.Unix(0, lTs*int64(time.Millisecond))
		dtm := time.Unix(0, dTs*int64(time.Millisecond))
		ctm := time.Unix(0, cTs*int64(time.Millisecond))
		switch platform {
			case "lion":
				fmt.Printf("\rLion Time: %s", ltm.Format("2006-01-02 15:04:05.000"))
			case "dog":
				fmt.Printf("\rDogTime: %s", dtm.Format("2006-01-02 15:04:05.000"))
			case "cat":
				fmt.Printf("\rCat Time: %s", ctm.Format("2006-01-02 15:04:05.000"))
			case "sys":
				fmt.Printf("\rSystem Time: %s", time.Now().Format("2006-01-02 15:04:05.000"))
			default:
				fmt.Printf("\rSystem Time: %s", time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func CatDeltaTime() int64 {
	cTs, _ := fetchTime("cat")
	sTs := time.Now().UnixNano() / 1e6  // system time
	delta := sTs - cTs
	// fmt.Println("cat delta:", delta)
	return delta
}

func DogDeltaTime() int64 {
	cTs, _ := fetchTime("dog")
	sTs := time.Now().UnixNano() / 1e6  // system time
	delta := sTs - cTs
	// fmt.Println("dog delta:", delta)
	return delta
}

func LionDeltaTime() int64 {
	cTs, _ := fetchTime("lion")
	sTs := time.Now().UnixNano() / 1e6  // system time
	delta := sTs - cTs
	// fmt.Println("lion delta:", delta)
	return delta
}

func LionDeltaTimeBySecond() int64 {
	uTs, _ := fetchTime("lion-s")
	fmt.Println("start", uTs)
	iTs := uTs
	count := 0
	var c1, c2, dCount int
	var lastTs, sTs, delta int64
	Loop:
		for {
			lastTs = uTs
			uTs, _ = fetchTime("lion-s")
			count ++
			if (lastTs != uTs) && (uTs == iTs + 1) {
				c1 = count
			} else if (lastTs != uTs) && (uTs == iTs + 2) {
				c2 = count
				sTs = time.Now().UnixNano() / 1e6  // system time
				dCount = c2 - c1
				requestSpendTime := 1000/dCount
				msUts := uTs*1000 + int64(requestSpendTime)  // added network spend time per request
				delta = sTs - msUts
				fmt.Println(sTs, msUts)
				break Loop
			}
			time.Sleep(100*time.Millisecond)
		}
	fmt.Println(iTs, uTs, sTs, c1, c2, delta)
	fmt.Println("lion-s delta:", delta)
	return delta
}



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

#lion-s's time
{
	"sysTime2": "2021-01-16 15:32:39",
	"sysTime1": "20210116153239"
}

#lion's time
{
	"api": "time",
	"code": "1",
	"currentTime": 1611300914188,
	"msg": ""
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
		datetimeStr := time.Unix(0, ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.000")
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	case "lion-s":  // LION_SECOND
		url = LION_SECOND_URL
		resp, _ := resty.New().R().Get(url)
		datetimeStr := gjson.Get(string(resp.Body()), "sysTime2").String()
		oTime, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.Local)
		ts := oTime.Unix()
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	case "lion":
		url = LION_URL
		resp, _ := resty.New().R().Get(url)
		ts := gjson.Get(string(resp.Body()), "currentTime").Int()
		datetimeStr := time.Unix(0, ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.000")
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	default:
		url = CAT_URL
		resp, _ := resty.New().R().Get(url)
		tsStr := gjson.Get(string(resp.Body()), "data.t").String()
		ts, _ := strconv.ParseInt(tsStr, 10, 64)
		datetimeStr := time.Unix(0, ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.000")
		// fmt.Println(ts)
		// fmt.Println(datetimeStr)
		return ts, datetimeStr
	}
}