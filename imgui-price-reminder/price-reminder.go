package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
	"image/color"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	board       [2]string
	defaultFont imgui.Font
	lcdFont     imgui.Font
	t           = 0
	posX        = 1000
	posY        = 50
	wnd         *g.MasterWindow
)

// PathExists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	return exePath
}

func initFont() {
	fonts := g.Context.IO().Fonts()
	ranges := imgui.NewGlyphRanges()
	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("0123456789./:(+) ")
	// builder.AddRanges(fonts.GlyphRangesChineseFull())
	builder.BuildRanges(ranges)
	/*
	defaultFont = fonts.AddFontFromFileTTFV("./ProggyTiny.ttf",
		10,
		imgui.DefaultFontConfig,
		fonts.GlyphRangesDefault())
	 */
	defaultFont = fonts.AddFontDefaultV(imgui.DefaultFontConfig)
	exePath := getExePath()
	lcdFontPath := exePath + string(os.PathSeparator) + "LcdD.ttf"
	if pathExisted, _ := PathExists(lcdFontPath); pathExisted {
		fmt.Println("lcdFontPath Existed")
		lcdFont = fonts.AddFontFromFileTTFV(
			lcdFontPath,
			50,
			imgui.DefaultFontConfig,
			ranges.Data())
	} else {
		lcdFont = fonts.AddFontFromFileTTFV(
			"./LcdD.ttf",
			50,
			imgui.DefaultFontConfig,
			ranges.Data())
	}
}

func onClickXAU() {
	openUrl("XAUUSD")
}

func onClickXAG() {
	openUrl("XAGUSD")
}

func onClickAUTD() {
	openUrl("Au(T+D)")
}

func onClickAGTD() {
	openUrl("Ag(T+D)")
}

func onClickUSDIDX() {
	openUrl("USDIDX")
}

func onChangeTheme() {
	t ++
	r := t%3
	switch r {
	case 0:
		imgui.StyleColorsDark()
		break
	case 1:
		imgui.StyleColorsClassic()
		break
	case 2:
		imgui.StyleColorsLight()
		break
	}
}

func openUrl(symbol string) {
	open.Run("https://chart.tubiaojia.com/tubiaojia.html?symbol=" + symbol)
}

func onExit() {
	os.Exit(0)
}

func refreshData() {
	count := 1
	ticker := time.NewTicker(time.Second * 15)
	for {
		count ++
		idx := count%4
		tips := [4]string{
			"XAUUSD\nusd/oz",
			"XAGUSD\nusd/oz",
			"Au(T+D)\ncny/g",
			"Ag(T+D)\ncny/kg",
		}
		price := fetchGoldAndSilverPrice(idx)
		if idx == 3 {
			price = strings.Replace(price, ".00", "", -1)
		}
		g.Update()
		tip := tips[idx]
		board = [2]string{price, tip}
		<- ticker.C
	}
}

func running() {
	layout := func() []g.Widget {
		var layout []g.Widget
		layout = g.Layout{
			g.MenuBar(g.Layout{
				g.Menu("Links", g.Layout{
					g.MenuItem("XAUUSD", onClickXAU),
					g.MenuItem("XAGUSD", onClickXAG),
					g.MenuItem("USDIDX", onClickUSDIDX),
					g.MenuItem("Au(T+D)", onClickAUTD),
					g.MenuItem("Ag(T+D)", onClickAGTD),
				}),
				g.Menu("Setting", g.Layout{
					g.MenuItem("ChangeTheme", onChangeTheme),
					g.MenuItem("Exit", onExit),
				}),
			}),
		}
		yellow := &color.RGBA{255, 255, 0, 255}
		layout = append(layout, g.Line(
			g.LabelV(fmt.Sprintf("%s", board[0]), false, yellow, &lcdFont),
			g.LabelV(board[1], false, nil, &defaultFont),
		))
		layout = append(layout, g.Label("(~^_^~) PriceReminder"))
		return layout
	}
	g.SingleWindowWithMenuBar("Price Reminder", layout())
}

func main() {
	exePath := getExePath()
	localEnvPath := exePath + string(os.PathSeparator) + ".env"
	if pathExisted, _ := PathExists(localEnvPath); pathExisted {
		godotenv.Load(localEnvPath)
	} else {
		godotenv.Load()
	}
	x := os.Getenv("PR_POS_X")
	y := os.Getenv("PR_POS_Y")
	posX, _ = strconv.Atoi(x)
	posY, _ = strconv.Atoi(y)
	if posX <= 0 {
		posX = 1000
	}
	if posY <= 0 {
		posY = 50
	}
	fmt.Println(posX, posY)
	flags := g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent
	// flags := g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsTransparent
	wnd = g.NewMasterWindow("Price Reminder", 180, 120, flags, initFont)
	bg := color.RGBA{}
	wnd.SetBgColor(bg)
	wnd.SetPos(posX, posY)
	go refreshData()
	wnd.Main(running)
}

func fetchGoldAndSilverPrice(idx int) string {
	// Sina 新浪源：`{code}` 目前支持  `hf_XAU` - XAU 国际黄金, `hf_XAG` - XAG 国际白银, `gds_AUTD` - AUTD 上海延期黄金, `gds_AGTD` - AGTD 上海延期白银
	symbols := []string{
		"hf_XAU",
		"hf_XAG",
		"gds_AUTD",
		"gds_AGTD",
	}
	providerUrl := "https://hq.sinajs.cn/?_={timestamp}/&list={code}"
	timestamp := time.Now().Unix()*1000 + int64(rand.Intn(899)+100)
	tsStr := strconv.FormatInt(timestamp, 10)
	code := symbols[idx]
	url := strings.Replace(providerUrl, "{timestamp}", tsStr, -1)
	url = strings.Replace(url, "{code}", code, -1)
	resp, err := resty.New().R().Get(url)
	if  err != nil {
		return "-"
	}
	content := strings.Replace(string(resp.Body()), `var hq_str_` + code + `="`, "", -1)
	content = strings.Replace(content, `";`, "", -1)
	result := strings.Split(content, ",")
	price := result[0]
	return price
}