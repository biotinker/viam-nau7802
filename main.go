// package main is a module providing a nau7802 driver
package main

import (
	"context"

	"github.com/biotinker/viam-nau7802/nau7802"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, logging.NewLogger("nau7802"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	nau7802Module, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	nau7802Module.AddModelFromRegistry(ctx, sensor.API, nau7802.Model)

	err = nau7802Module.Start(ctx)
	defer nau7802Module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

