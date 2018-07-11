package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Microsoft/hcsshim/internal/guid"
	"github.com/Microsoft/hcsshim/internal/hcs"
	"github.com/Microsoft/hcsshim/internal/schema2"
	"github.com/Microsoft/hcsshim/internal/schemaversion"

	"github.com/sirupsen/logrus"
)

const (
	timeoutSeconds = time.Second * 60 * 4
)

func main() {

	// Set some variables. Could come from command line options
	id := guid.New().String()
	owner := filepath.Base(os.Args[0])

	memory := int32(2 * 1024) // in MB -> so 2 GB
	processors := int32(2)

	bootFilePath := `C:\Program Files\Linux Containers`
	kernelFile := "bootx64.efi"
	initrdFile := "initrd.img"
	kernelArgs := "initrd=/" + initrdFile + " page_poison=1 vsyscall=emulate panic=1"

	consolePipe := `\\localhost\pipe\vmpipe`

	scsi := make(map[string]hcsschema.Scsi)

	// Basic VM setup. We add more stuff below.
	vm := &hcsschema.VirtualMachine{
		Chipset: &hcsschema.Chipset{
			Uefi: &hcsschema.Uefi{},
		},

		ComputeTopology: &hcsschema.Topology{
			Memory: &hcsschema.Memory2{
				Backing:         "Virtual",
				InitialSizeInMB: memory,
			},
			Processor: &hcsschema.Processor2{
				Count: processors,
			},
		},

		Devices: &hcsschema.Devices{
			Scsi: scsi,
		},
	}

	// Serial console
	kernelArgs += " console=ttyS0,115200"
	vm.Devices.ComPorts = map[string]hcsschema.ComPort{
		"0": { // Which is actually COM1
			NamedPipe: consolePipe,
		},
	}

	// Specify where the boot files are.
	// Looks like they are passed as a Virtual SMB Share (whatever that is)
	vm.Devices.VirtualSmbShares = []hcsschema.VirtualSmbShare{
		{
			Name: "os",
			Path: bootFilePath,
			Options: &hcsschema.VirtualSmbShareOptions{
				ReadOnly:            true,
				TakeBackupPrivilege: true,
				CacheIo:             true,
				ShareRead:           true,
			},
		},
	}

	vm.Chipset.Uefi.BootThis = &hcsschema.UefiBootEntry{
		DevicePath:   `\` + kernelFile,
		DeviceType:   "VMBFS",
		OptionalData: kernelArgs,
	}

	hcsDocument := &hcsschema.ComputeSystem{
		Owner:          owner,
		SchemaVersion:  schemaversion.SchemaV20(),
		VirtualMachine: vm,
	}

	hcsSystem, err := hcs.CreateComputeSystem(id, hcsDocument)
	if err != nil {
		logrus.Fatalln("failed to create VM: ", err)
	}

	logrus.Infoln("Created: %v", hcsSystem)
}
