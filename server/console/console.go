package console

import (
	"fmt"
	ui "github.com/gizak/termui"
	"rigpig/internal"
	"rigpig/internal/common"
	"strconv"
	"time"
)

var Agents []Agent

type Agent struct {
	Address    string
	Status     string
	HashRate   string
	Accepted   string
	Rejected   string
	Algo       string
	Pool       string
	Miner      string
	Uptime     string
	MiningTime string
	Coin       string
	Profit     string
}

var CoinData []internal.AlgoStats

func MakeConsole(topAlgoStats <-chan []internal.AlgoStats) {

	Agents = append(Agents, Agent{Address: "192.168.1.20", Status: "Online", HashRate: "45.3", Accepted: "340", Rejected: "2", Algo: "x16r", Pool: "pw.bsod.org:3326", Miner: "ccminer", Uptime: "00:10:03:00", Coin: "Unspecified", Profit: "0.00"})
	Agents = append(Agents, Agent{Address: "172.16.10.3", Status: "Online", HashRate: "45.3", Accepted: "340", Rejected: "2", Algo: "x16r", Pool: "pw.bsod.org:4250", Miner: "Enemy 1.14a", Uptime: "00:10:04:24", Coin: "Ravencoin", Profit: "3.49"})
	Agents = append(Agents, Agent{Address: "10.0.0.254", Status: "Online", HashRate: "45.3", Accepted: "340", Rejected: "2", Algo: "x16r", Pool: "pw.bsod.org:3326", Miner: "ccminer", Uptime: "00:05:45:23", Coin: "Ravencoin", Profit: "2.43"})
	Agents = append(Agents, Agent{Address: "192.168.1.50", Status: "Online", HashRate: "45.3", Accepted: "340", Rejected: "2", Algo: "x16r", Pool: "pw.bsod.org:3326", Miner: "ccminer", Uptime: "00:17:20:56", Coin: "PigeonCoin", Profit: "4.50"})
	Agents = append(Agents, Agent{Address: "192.168.1.51", Status: "Online", HashRate: "45.3", Accepted: "340", Rejected: "2", Algo: "x16r", Pool: "pw.bsod.org:4250", Miner: "z-Enemy 1.12", Uptime: "00:02:25:23", Coin: "PigeonCoin", Profit: "4.23"})

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar(common.ApplicationName)
	p.Height = 1
	p.Width = 100
	p.TextFgColor = ui.ColorWhite
	p.BorderBottom = false
	p.BorderTop = false
	p.BorderLeft = false
	p.BorderRight = false
	p.X = 1
	p.Y = 0

	subtitle := ui.NewPar(common.ApplicationSubtitle)
	subtitle.Height = 1
	subtitle.Width = 100
	subtitle.TextFgColor = ui.ColorYellow
	subtitle.Border = false
	subtitle.X = 1
	subtitle.Y = 1

	s1 := ui.NewPar("API SERVER\n" +
		"REMOTE AGENT SERVER\n" +
		"WEB CONSOLE")
	s1.Height = 10
	s1.Width = 40
	s1.TextFgColor = ui.ColorWhite
	s1.Border = false
	s1.X = 1
	s1.Y = 4
	s1.WrapLength = 40

	s1Status := ui.NewPar("RUNNING    3000\n" +
		"RUNNING    3001\n" +
		"RUNNING    8080")
	s1Status.Height = 10
	s1Status.Width = 40
	s1Status.TextFgColor = ui.ColorCyan
	s1Status.Border = false
	s1Status.X = 30
	s1Status.Y = 4
	s1Status.WrapLength = 40

	/*
		lc := ui.NewLineChart()
		sinps := (func() []float64 {

			ps := make([]float64,1000)
			for i := range ps {
				ps[i] = float64(i) * rand.Float64()
			}
			return ps
		})()

		lc.BorderLabel = "COMBINED HASHRATE"
		lc.Data = sinps
		lc.Width = 120
		lc.Height = 25
		lc.LineColor = ui.ColorCyan
		lc.X = 1
		lc.Y = 7
		//lc.Mode = "dot"
	*/

	rowst1 := [][]string{}
	rowst1 = append(rowst1, []string{"RIG", "PROFILE", "STATUS", "UPTIME", "LAST UPDATE", "HASHRATE", "ACCEPT", "REJECT", "ALGO", "MINER", "COIN", "POOL", "PROFIT"})

	for _, v := range Agents {
		lastUpdateMock := fmt.Sprintf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		row := []string{v.Address, "NVIDIA GPU", v.Status, v.Uptime, lastUpdateMock, v.HashRate, v.Accepted, v.Rejected, v.Algo, v.Miner, v.Coin, v.Pool, v.Profit}
		rowst1 = append(rowst1, row)
	}

	t1 := ui.NewTable()
	t1.Rows = RigTableRows()
	t1.PaddingTop = 0
	t1.PaddingLeft = -2
	t1.FgColor = ui.ColorWhite
	t1.BgColor = ui.ColorDefault
	t1.BorderLabel = "MINING RIGS"
	t1.Border = true
	t1.BorderFg = ui.ColorBlack
	t1.BorderBg = ui.ColorBlack
	t1.Separator = false
	t1.Analysis()
	t1.SetSize()
	t1.FgColors[0] = ui.ColorCyan
	t1.Y = 12
	t1.X = 0

	//CoinData = internal.GetLatestAlgoStats()

	coinsTable := ui.NewTable()
	coinsTable.Rows = CreateCoinDataTable()
	coinsTable.PaddingTop = 0
	coinsTable.PaddingLeft = -2
	coinsTable.BorderLabel = "TOP POOL ALGOS"
	coinsTable.Border = true
	coinsTable.BorderFg = ui.ColorBlack
	coinsTable.BorderBg = ui.ColorBlack
	coinsTable.Separator = false
	coinsTable.Analysis()
	coinsTable.SetSize()
	coinsTable.Y = 21
	coinsTable.X = 0

	poolStatus := ui.NewTable()
	poolStatus.Rows = CreatePoolStatusRows()
	poolStatus.PaddingTop = 0
	poolStatus.PaddingLeft = -2
	poolStatus.BorderLabel = "MINING POOLS"
	poolStatus.Border = true
	poolStatus.BorderFg = ui.ColorBlack
	poolStatus.BorderBg = ui.ColorBlack
	poolStatus.Separator = false
	poolStatus.Analysis()
	poolStatus.SetSize()
	poolStatus.Y = 21
	poolStatus.X = 100
	poolStatus.FgColors[0] = ui.ColorCyan

	draw := func(t int) {
		//lc.Data = sinps[t/2:]
		coinsTable.Rows = CreateCoinDataTable()
		coinsTable.FgColors[0] = ui.ColorCyan
		coinsTable.Analysis()
		coinsTable.SetSize()
		coinsTable.BorderLabel = "20 MOST PROFITABLE ALGOS -- " + fmt.Sprintf("UPDATED: %02d:%02d:%02d", internal.Updated.Hour(), internal.Updated.Minute(), internal.Updated.Second())

		poolStatus.Rows = CreatePoolStatusRows()
		t1.Rows = RigTableRows()

		ui.Render(p, subtitle, s1, s1Status, t1, coinsTable, poolStatus)
	}

	ui.Render(p, subtitle, s1, s1Status, t1, coinsTable, poolStatus)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})

	ui.Loop()
}

func RigTableRows() [][]string {
	rowst1 := [][]string{}
	rowst1 = append(rowst1, []string{"RIG", "PROFILE", "STATUS", "UPTIME", "LAST UPDATE", "HASHRATE", "ACCEPT", "REJECT", "ALGO", "MINER", "COIN", "POOL", "PROFIT"})

	for _, v := range Agents {
		lastUpdateMock := fmt.Sprintf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		row := []string{v.Address, "NVIDIA GPU", v.Status, v.Uptime, lastUpdateMock, v.HashRate, v.Accepted, v.Rejected, v.Algo, v.Miner, v.Coin, v.Pool, v.Profit}
		rowst1 = append(rowst1, row)
	}
	return rowst1

}

func CreatePoolStatusRows() [][]string {
	rows := [][]string{}
	rows = append(rows, []string{"POOL", "STATUS", "LAST UPDATE"})

	if len(internal.Pools) > 0 {
		for pool := range internal.Pools {
			var status string
			if internal.Pools[pool].Enabled == true {
				status = "Enabled"
			} else {
				status = "Disabled"
			}
			lastUpdate := fmt.Sprintf("%02d:%02d:%02d", internal.Pools[pool].LastUpdate.Hour(), internal.Pools[pool].LastUpdate.Minute(), internal.Pools[pool].LastUpdate.Second())
			rows = append(rows, []string{internal.Pools[pool].Name, status, lastUpdate})
		}
	}

	return rows
}

func CreateCoinDataTable() [][]string {

	CoinData = internal.GetLatestAlgoStats()
	coinsTableRows := [][]string{}
	coinsTableRows = append(coinsTableRows, []string{"POS", "ALGO", "POOL", "PORT", "BTC ESTIMATE 24h", "BTC ACTUAL 24h", "POOL WORKERS"})

	if len(CoinData) > 0 {
		for i := 0; i <= 39; i++ {
			coinsTableRows = append(coinsTableRows, []string{strconv.Itoa(i + 1), CoinData[i].Name, CoinData[i].PoolName, strconv.Itoa(CoinData[i].Port), CoinData[i].EstimateLast24h, CoinData[i].ActualLast24h, strconv.Itoa(CoinData[i].Workers)})
		}
	} else {

		for i := 0; i <= 39; i++ {
			coinsTableRows = append(coinsTableRows, []string{"", "", "", "", ""})
		}

	}

	return coinsTableRows
}
