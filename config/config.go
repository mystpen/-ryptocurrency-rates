package config

import "time"

var Config  = struct {
	Url      string
	Interval time.Duration
}{
	Url: "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1",
	Interval: 10 * time.Minute,
}