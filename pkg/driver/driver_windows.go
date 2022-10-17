package driver

import (
	"github.com/StackExchange/wmi"
	"strings"
)

const wqlPnPSignedDriver = "SELECT DeviceClass, DeviceID, DeviceName, DriverDate, DriverName, DriverVersion, FriendlyName, HardWareID, Name, DriverProviderName FROM Win32_PnPSignedDriver"

type win32PnPSignedDriver struct {
	DeviceClass        string
	DeviceID           string
	DeviceName         string
	DriverDate         string
	DriverName         string
	DriverVersion      string
	FriendlyName       string
	HardWareID         string
	Name               string
	DriverProviderName string
}

func (i *Info) load() error {
	// Getting data from WMI
	var win32PnPSignedDriverDescriptions []win32PnPSignedDriver
	if err := wmi.Query(wqlPnPSignedDriver, &win32PnPSignedDriverDescriptions); err != nil {
		return err
	}

	// Building dynamic WHERE clause with addresses to create a single query collecting all desired data
	queryAddresses := []string{}
	for _, description := range win32PnPSignedDriverDescriptions {
		var queryAddress = strings.Replace(description.DeviceID, "\\", `\\`, -1)
		queryAddresses = append(queryAddresses, "PNPDeviceID='"+queryAddress+"'")
	}

	// Converting into standard structures
	drivers := make([]*Driver, 0)
	for _, description := range win32PnPSignedDriverDescriptions {
		driver := &Driver{
			Address: description.DeviceID,
			Index:   0,
		}
		drivers = append(drivers, driver)
	}
	i.Drivers = drivers
	return nil
}
