package main

import (
	"os"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"github.com/ebiiim/toy-plugins/pkg/maintainer"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(maintainer.Name, maintainer.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()
	
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
