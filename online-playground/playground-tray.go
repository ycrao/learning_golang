package main

import (
	"encoding/json"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"math/rand"
	"online-playground/icon"
	"os"
	"time"
)

func main() {
	systray.Run(onReady, nil)
}


type MenuItems struct {
	MenuItems []MenuItem `json:"menu_items"`
}

type MenuItem struct {
	Name string `json:"name"`
	Urls []string `json:"urls"`
}

func onReady() {
	jsonFile, err := os.Open("playground.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	playgroundByteValue, _ := ioutil.ReadAll(jsonFile)
	var menuItems MenuItems
	json.Unmarshal(playgroundByteValue, &menuItems)
	iconData := icon.Data
	menuItemsLen := len(menuItems.MenuItems)
	go func() {
		systray.SetTemplateIcon(iconData, iconData)
		// systray.SetTitle("Playground")
		systray.SetTooltip("Get Online Playground by one-click!")
		for i := 0; i < menuItemsLen; i ++ {
			name := menuItems.MenuItems[i].Name
			urls := menuItems.MenuItems[i].Urls
			tempItem := systray.AddMenuItem(name, "Open " + name + "Playground")
			go func() {
				for {
					select {
					case <-tempItem.ClickedCh:
						tLen := len(urls)
						rand.Seed(int64(time.Now().UnixNano()))
						idx := rand.Intn(99)%tLen
						open.Run(urls[idx])
					}
				}
			}()
		}
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit the App")
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}