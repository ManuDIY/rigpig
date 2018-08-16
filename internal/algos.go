package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"time"
)

var (
	pools []MiningPool
)

// MinedAlgos stores a list of algos mined
type MinedAlgos struct {
	Algos []Algo
}

// Algo contains a collection for all supported algos used to mine
// We hard code these values to avoid security issues by importing
// json as interface{}.
type Algo struct {
	Aergo       AlgoStats `json:"aergo"`
	Allium      AlgoStats `json:"allium"`
	Ballon      AlgoStats `json:"ballon"`
	Bitcore     AlgoStats `json:"bitcore"`
	Blake2s     AlgoStats `json:"blake2s"`
	Blakecoin   AlgoStats `json:"blakecoin"`
	C11         AlgoStats `json:"c11"`
	Equihash    AlgoStats `json:"equihash"`
	Equihash144 AlgoStats `json:"equihash144"`
	Equihash192 AlgoStats `json:"equihash192"`
	Groestl     AlgoStats `json:"groestl"`
	Hex         AlgoStats `json:"hex"`
	Hmq1725     AlgoStats `json:"hmq1725"`
	Keccak      AlgoStats `json:"keccak"`
	Keccakc     AlgoStats `json:"keccakc"`
	Lbry        AlgoStats `json:"lbry"`
	Lyra2v2     AlgoStats `json:"lyra2v2"`
	Lyra2z      AlgoStats `json:"lyra2z"`
	M7m         AlgoStats `json:"m7m"`
	MyrGr       AlgoStats `json:"myr-gr"`
	Neoscript   AlgoStats `json:"neoscript"`
	Nist5       AlgoStats `json:"nist5"`
	Phi         AlgoStats `json:"phi"`
	Phi2        AlgoStats `json:"phi2"`
	Quark       AlgoStats `json:"quark"`
	Qubit       AlgoStats `json:"qubit"`
	Scrypt      AlgoStats `json:"scrypt"`
	Sib         AlgoStats `json:"sib"`
	Sha256      AlgoStats `json:"sha256"`
	Skein       AlgoStats `json:"skein"`
	Skunk       AlgoStats `json:"skunk"`
	TimeTravel  AlgoStats `json:"timetravel"`
	Tribus      AlgoStats `json:"tribus"`
	X11         AlgoStats `json:"x11"`
	X16r        AlgoStats `json:"x16r"`
	X16s        AlgoStats `json:"x16s"`
	X17         AlgoStats `json:"x17"`
	Xevan       AlgoStats `json:"xevan"`
	Yescrypt    AlgoStats `json:"yescrypt"`
	YescryptR16 AlgoStats `json:"yescryptR16"`
	YesPower    AlgoStats `json:"yespower"`
}

// AlgoStats stores statistics of a pool's algo
type AlgoStats struct {
	Name                              string  `json:"name"`
	Port                              int     `json:"port"`
	Coins                             int     `json:"coins"`
	Fees                              float32 `json:"fees"`
	Hashrate                          float64 `json:"hashrate"`
	Workers                           int     `json:"workers"`
	EstimateCurrent                   string  `json:"estimate_current"`
	EstimateLast24h                   string  `json:"estimate_last24h"`
	ActualLast24h                     string  `json:"actual_last24h"`
	MbtcMhFactor                      float64 `json:"mbtc_mh_factor"`
	HashrateLast24h                   float64 `json:"hashrate_last24h"`
	ActualLast24hInBtcPerHashPerDay   string  `json:"actual_last24h_in_btc_per_has_per_day"`
	EstimateCurrentInBtcPerHashPerDay string  `json:"estimate_current_in_btc_per_hash_per_day"`
	EstimateLast24hInBtcPerHashPerDay string  `json:"estimate_last24h_in_btc_per_hash_per_day"`
	PoolName                          string
	PoolURL                           string
}

type MiningPool struct {
	Name           string
	URL            string
	PayoutCurrency string
	DefaultUnits   string
	Enabled        bool
	LastUpdate     time.Time
}

var totalAlgos int
var TopAlgos []AlgoStats
var OutputAlgoStats []AlgoStats
var Updated time.Time
var Pools []MiningPool

func UpdateAlgos() []AlgoStats {

	// Split this out in the future. Put logic to enabled\disable pools
	Pools = append(Pools, MiningPool{Name: "ahashpool", URL: "https://www.ahashpool.com/api/status", DefaultUnits: "BTC", Enabled: true})
	Pools = append(Pools, MiningPool{Name: "blazepool", URL: "http://api.blazepool.com/status", DefaultUnits: "BTC", Enabled: true})
	//Pools = append(Pools, MiningPool{ Name: "hashrefinery",	URL: "http://pool.hashrefinery.com/api/status",	DefaultUnits: "BTC", Enabled: false})
	Pools = append(Pools, MiningPool{Name: "zergpool", URL: "http://api.zergpool.com:8080/api/status", DefaultUnits: "BTC", Enabled: true})
	Pools = append(Pools, MiningPool{Name: "zpool", URL: "http://www.zpool.ca/api/status", DefaultUnits: "BTC", Enabled: true})

	pool_stats, _ := fetchCombinedPools()
	combinedAlgos := findMostProfitable(pool_stats)

	// Create a slice that will sort combinedAlgos by key
	//log.Printf("Crunching and ordering by profitability...\n")
	var keys []float64
	for k := range combinedAlgos {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	// Create an ordered list of algos, sorted by most profitable first
	totalAlgos = len(keys) - 1

	if totalAlgos < 0 {
		totalAlgos = 0
	}
	TopAlgos = make([]AlgoStats, totalAlgos)

	//log.Print("Building sorted list: sorting by most profitable first")
	listPosition := 0
	for i := totalAlgos; i >= 1; i-- {
		TopAlgos[listPosition] = combinedAlgos[keys[i]]
		listPosition++
	}

	Updated = time.Now()
	return TopAlgos
}

func GetLatestAlgoStats() []AlgoStats {
	return OutputAlgoStats
}

func fetchCombinedPools() (map[MiningPool]Algo, error) {
	// Set request timeout
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	pool_stats := make(map[MiningPool]Algo, 500)

	//log.Printf("Collection lastest crypto stats:")

	// Iterate through all pools and
	for k, p := range Pools {

		// log.Printf("==> Hitting : %s", p.Name)

		var algos Algo

		resp, err := httpClient.Get(p.URL)
		if err != nil {
			return pool_stats, err
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return pool_stats, err
		}

		err = json.Unmarshal(bytes, &algos)
		if err != nil {
			return pool_stats, err
		}
		pool_stats[p] = algos

		Pools[k].LastUpdate = time.Now()

	}
	return pool_stats, nil
}

func findMostProfitable(pool_stats map[MiningPool]Algo) (combinedAlgos map[float64]AlgoStats) {
	combinedAlgos = make(map[float64]AlgoStats)

	for pool, pstat := range pool_stats {
		v := reflect.ValueOf(pstat)

		for i := 0; i < v.NumField(); i++ {
			vs := reflect.ValueOf(v.Field(i).Interface())
			current24 := vs.FieldByName("ActualLast24h")

			//fmt.Println(vs.FieldByName("Name").String(), current24.String())

			if current24.String() != "" {
				current24Float, err := strconv.ParseFloat(current24.String(), 64)
				if err != nil {
					log.Fatal(err, current24)
				}
				combinedAlgos[current24Float] = AlgoStats{
					Name:            v.Field(i).FieldByName("Name").String(),
					EstimateLast24h: v.Field(i).FieldByName("EstimateLast24h").String(),
					ActualLast24h:   v.Field(i).FieldByName("ActualLast24h").String(),
					Port:            int(v.Field(i).FieldByName("Port").Int()),
					Hashrate:        float64(v.Field(i).FieldByName("Hashrate").Float()),
					Workers:         int(v.Field(i).FieldByName("Workers").Int()),
					PoolName:        pool.Name,
					PoolURL:         pool.URL,
				}
			}
		}
	}

	return
}

// DEPRECATED
// Replace with version that supports multiple pools
func sortAlgoByProfit(algos Algo) map[float64]AlgoStats {
	sorted := make(map[float64]AlgoStats)

	v := reflect.ValueOf(algos)

	//values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		vs := reflect.ValueOf(v.Field(i).Interface())
		current24 := vs.FieldByName("EstimateLast24h")

		//fmt.Println(vs.FieldByName("Name").String(), current24.String())

		if current24.String() != "" {
			current24Float, err := strconv.ParseFloat(current24.String(), 64)
			if err != nil {
				log.Fatal(err, current24)
			}

			sorted[current24Float] = AlgoStats{
				Name:            v.Field(i).FieldByName("Name").String(),
				EstimateLast24h: v.Field(i).FieldByName("EstimateLast24h").String(),
				Port:            int(v.Field(i).FieldByName("Port").Int()),
				Hashrate:        float64(v.Field(i).FieldByName("Hashrate").Float()),
			}
		}

		// values[i] = v.Field(i).Interface()
		//sorted[v.Field(1).FieldByName("EstimateCurrent")] = v.Field(1)
		//typeOfT.Fi
	}

	//fmt.Println(values)

	var keys []float64
	for k := range sorted {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	highestAlgo := len(keys) - 1

	for i := highestAlgo; i > 0; i-- {
		fmt.Printf("Name: %s\tLast24h: %s\tPort: %d\tHashrate: %0.2f\n", sorted[keys[i]].Name, sorted[keys[i]].EstimateLast24h, sorted[keys[i]].Port, sorted[keys[i]].Hashrate)
	}

	//fmt.Printf("%0.8f\n", sorted[keys[highestAlgo]])

	/*
		for _, k := range keys {
			//fmt.Printf("%s", sorted[k])
			fmt.Printf("%0.8f\n", sorted[highestAlgo])
			//fmt.Println(sorted[k])
		}
	*/
	return sorted
}
