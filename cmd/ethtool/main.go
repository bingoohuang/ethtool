package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bingoohuang/ethtool"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("interface is not specified")
	}

	e, err := ethtool.NewEthtool()
	if err != nil {
		panic(err.Error())
	}
	defer e.Close()

	name := os.Args[1]

	if features, err := e.Features(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("features: %+v\n", features)
	}

	if stats, err := e.Stats(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("stats: %+v\n", stats)
	}

	if busInfo, err := e.BusInfo(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("bus info: %+v\n", busInfo)
	}

	if drvr, err := e.DriverName(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("driver name: %+v\n", drvr)
	}

	if cmdGet, err := e.CmdGetMapped(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("cmd get: %+v\n", cmdGet)
	}

	if msgLvlGet, err := e.MsglvlGet(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("msg lvl get: %+v\n", msgLvlGet)
	}

	if drvInfo, err := e.DriverInfo(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("drvrinfo: %+v\n", drvInfo)
	}

	if permAddr, err := e.PermAddr(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("permaddr: %+v\n", permAddr)
	}

	if eeprom, err := e.ModuleEepromHex(name); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("module eeprom: %+v\n", eeprom)
	}
}
