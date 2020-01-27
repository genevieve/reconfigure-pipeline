package main

import (
	"os"

	"github.com/genevieve/reconfigure-pipeline/actions"
	"github.com/genevieve/reconfigure-pipeline/commandrunner"
	"github.com/genevieve/reconfigure-pipeline/concourse"
	"github.com/genevieve/reconfigure-pipeline/lastpass"
	"github.com/genevieve/reconfigure-pipeline/writer"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		ConfigPath     string `short:"c" long:"config" required:"true" description:"Pipeline configuration file"`
		Pipeline       string `short:"p" long:"pipeline" required:"true" description:"Pipeline to configure"`
		Target         string `short:"t" long:"target" required:"true" description:"Concourse target name"`
		VariablesPath  string `short:"l" long:"load-vars-from" description:"Variable flag that can be used for filling in template values in configuration from a YAML file"`
		NonInteractive bool   `short:"n" long:"non-interactive" description:"Skip confirmation"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}

	action := newReconfigurePipeline()
	err = action.Run(opts.Target, opts.Pipeline, opts.ConfigPath, opts.VariablesPath, opts.NonInteractive)

	if err != nil {
		os.Exit(1)
	}
}

func newReconfigurePipeline() *actions.ReconfigurePipeline {
	commandRunner := commandrunner.NewSimpleCommandRunner()

	reconfigurer := concourse.NewReconfigurer(commandRunner)
	processor := lastpass.NewProcessor(commandRunner)
	configWriter := writer.NewConfigWriter()

	return actions.NewReconfigurePipeline(reconfigurer, processor, configWriter)
}
