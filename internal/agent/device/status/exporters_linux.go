//go:build linux

package status

import (
	"os"

	"github.com/flightctl/flightctl/internal/agent/device/resource"
	"github.com/flightctl/flightctl/pkg/executer"
	"github.com/flightctl/flightctl/pkg/log"
)

func newExporters(
	resourceManager resource.Manager,
	executer executer.Executer,
	log *log.PrefixLogger,
) []Exporter {
	return []Exporter{
		newSystemD(executer),
		newContainer(executer),
		newSystemInfo(executer),
		newResources(log, resourceManager),
	}
}

func getBootID(bootIDPath string) (string, error) {
	bootID, err := os.ReadFile(bootIDPath)
	if err != nil {
		return "", err
	}
	return string(bootID), nil
}
