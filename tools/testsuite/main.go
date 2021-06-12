package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

type config struct {
	API struct {
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
	}
}

func configFromFile(path string) (c config, err error) {
	// Load file
	tree, err := toml.LoadFile(path)
	if err != nil {
		return config{}, err
	}

	// Change into structure
	err = tree.Unmarshal(&c)
	return c, err
}

func run() int {
	fmt.Println("Running Binance Integration TestSuite")

	conf, err := configFromFile("configs/testsuite.toml")
	if err != nil {
		fmt.Println("Error when reading the configuration:", err)
		return 255
	}

	count := runCandlestickTests(conf.API.Key, conf.API.Secret)

	fmt.Println("Testsuite finished")
	return count
}

func main() {
	os.Exit(run())
}
