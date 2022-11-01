package server

import (
	"math/big"
)

type Config struct {
	network    string
	httpPort   int
	interval   int
	payout     *big.Int
	proxyCount int
	queueCap   int
}

func NewConfig(network string, httpPort, interval, payout, proxyCount, queueCap int) *Config {
	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(17), nil)
	return &Config{
		network:    network,
		httpPort:   httpPort,
		interval:   interval,
		payout:     new(big.Int).Mul(big.NewInt(int64(payout)), ether),
		proxyCount: proxyCount,
		queueCap:   queueCap,
	}
}
