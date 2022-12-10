package main

import (
	"log"
	"math"
	"tinygo.org/x/bluetooth"
)

func main() {
	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	if err != nil {
		panic("failed to enable BLE adapter")
	}

	err = adapter.Scan(onScan)
	if err != nil {
		panic("failed to register scan callback")
	}
}

const BleServiceTemperature = 0xFFE0
const BleCharacteristicTemperature = 0xFFE1

func onScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	//log.Println("found device:", device.Address.String(), device.RSSI, device.LocalName(), device.AdvertisementPayload)
	if device.HasServiceUUID(bluetooth.New16BitUUID(BleServiceTemperature)) {
		log.Println("found device with temperature service:", device.Address.String(), device.RSSI, device.LocalName())
		go func() {
			res, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})
			if err != nil {
				log.Println("error connecting:", err.Error())
				return
			}
			onConnect(device, res)
		}()
	}
}

func onConnect(scanResult bluetooth.ScanResult, device *bluetooth.Device) {
	log.Println("connected:", scanResult.Address.String(), scanResult.LocalName())

	// Get a list of services
	services, err := device.DiscoverServices([]bluetooth.UUID{
		bluetooth.New16BitUUID(BleServiceTemperature),
	})
	// If error, bail out
	if err != nil {
		log.Println("error getting services:", err.Error())
		return
	}
	// Iterate services
	for _, service := range services {
		if service.UUID() != bluetooth.New16BitUUID(BleServiceTemperature) {
			// Wrong service
			continue
		}
		// Found the correct service
		// Get a list of characteristics below the service
		characteristics, err := service.DiscoverCharacteristics([]bluetooth.UUID{
			bluetooth.New16BitUUID(BleCharacteristicTemperature),
		})
		// If error, bail out
		if err != nil {
			log.Println("error getting characteristics:", err.Error())
			return
		}
		// Iterate characteristics
		for _, characteristic := range characteristics {
			err := characteristic.EnableNotifications(characteristicReceiverTemperature)
			// If error, bail out
			if err != nil {
				log.Println("error enabling notifications:", err.Error())
				return
			}
		}
	}
}

func characteristicReceiverTemperature(buf []byte) {
	log.Printf("received: %x", buf)

	if !isTemperatureSuccess(buf) {
		log.Printf("wrong measurement")
		return
	}

	temp := (float64(buf[4]) * 256) + float64(buf[5])
	temp = math.Round(temp/10) / 10
	log.Printf("temp: %v", temp)
}

var TemperatureMeasureSuccess = [...]byte{0xAA, 0x01, 0xC1}
var TemperatureMeasureError = [...]byte{0xAA, 0x01, 0xCE}

func isTemperatureSuccess(buf []byte) bool {
	if len(buf) < len(TemperatureMeasureSuccess) {
		return false
	}
	for i, b := range TemperatureMeasureSuccess {
		if buf[i] != b {
			return false
		}
	}
	return true
}
