package main

import (
	"fmt"

	"github.com/soupstore/coda-world/config"
	"github.com/soupstore/coda-world/data"
	"github.com/soupstore/coda-world/log"
	"github.com/soupstore/coda-world/simulation"
	"github.com/soupstore/coda-world/telnet"
)

func main() {
	log.Logger().Info("Starting world server")

	// load config values from env vars
	conf, err := config.Load()
	if err != nil {
		log.Logger().Fatal(err.Error())
	}

	// create a new simulation
	sim := simulation.NewSimulation()

	// load static data and apply changes to simulation
	dw := data.NewDataWatcher(conf.DataPath, sim)
	log.SubscribeToErrorChan(dw.Errors)

	// temporary
	sim.SetSpawnRoom("admin", 1)

	_ = sim.MakeCharacter("Rinse")
	sim.MakeCharacter("Claw")
	sim.MakeCharacter("Gesau")

	nexus, _ := sim.GetRoom("admin", 1)
	sim.SpawnItem(1, nexus.Container.ID)
	sim.SpawnItem(2, nexus.Container.ID)

	listenAddr := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	telnetServer := telnet.NewServer(conf, listenAddr, sim)
	err = telnetServer.ListenAndServe()
}
