# Bluetooth LE cadence and speed sensor example

Fork of https://github.com/Ecostack/go-ble-thermometer

GATT characteristics: https://www.bluetooth.com/de/specifications/gss

    [BK6LC-0056626]# list-attributes
    Primary Service (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025
        0000180a-0000-1000-8000-00805f9b34fb
        Device Information
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002c
        00002a28-0000-1000-8000-00805f9b34fb
        Software Revision String
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002a
        00002a27-0000-1000-8000-00805f9b34fb
        Hardware Revision String
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char0028
        00002a24-0000-1000-8000-00805f9b34fb
        Model Number String
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char0026
        00002a29-0000-1000-8000-00805f9b34fb
        Manufacturer Name String
    .
    .
    .
    Descriptor (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service001a/char001b/desc001d
        00002902-0000-1000-8000-00805f9b34fb
        Client Characteristic Configuration
    Primary Service (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016
        0000180f-0000-1000-8000-00805f9b34fb
        Battery Service
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016/char0017
        00002a19-0000-1000-8000-00805f9b34fb
        Battery Level
    Descriptor (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016/char0017/desc0019
        00002902-0000-1000-8000-00805f9b34fb
        Client Characteristic Configuration
    Primary Service (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b
        00001816-0000-1000-8000-00805f9b34fb
        Cycling Speed and Cadence
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char0013
        00002a55-0000-1000-8000-00805f9b34fb
        SC Control Point
    Descriptor (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char0013/desc0015
        00002902-0000-1000-8000-00805f9b34fb
        Client Characteristic Configuration
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char0011
        00002a5d-0000-1000-8000-00805f9b34fb
        Sensor Location
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000f
        00002a5c-0000-1000-8000-00805f9b34fb
        CSC Feature
    Characteristic (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c
        00002a5b-0000-1000-8000-00805f9b34fb
        CSC Measurement
    Descriptor (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c/desc000e
        00002902-0000-1000-8000-00805f9b34fb
        Client Characteristic Configuration
    Primary Service (Handle 0x0000)
        /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000a
        00001801-0000-1000-8000-00805f9b34fb
        Generic Attribute Profile

## Attribute Software Revision String

    [BK6LC-0056626:/service0025/char002c]# read
    Attempting to read /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002c
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002c Value:
    56 31 2e 34 2e 30                                V1.4.0

## Attribute Hardware Revision String

    [BK6LC-0056626:/service0025]# select-attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002a
    [BK6LC-0056626:/service0025/char002a]# read
    Attempting to read /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002a
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0025/char002a Value:
    56 33 2e 30 2e 30 4c                             V3.0.0L

## Attribute CSC Feature

    [BK6LC-0056626:/service000b/char0013]# select-attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000f
    [BK6LC-0056626:/service000b/char000f]# read
    Attempting to read /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000f
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000f Value:
    02 00                                            ..

    |01234567|01234567|
    +--------+--------+
    |01000000|00000000|

## Attribute Battery Level

    [BK6LC-0056626:/service0016/char0017]# select-attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016/char0017
    [BK6LC-0056626:/service0016/char0017]# read
    Attempting to read /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016/char0017
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service0016/char0017 Value:
    64                                               d

## Attribute CSC Measurement

    [BK6LC-0056626:/service0016/char0017]# select-attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c
    [BK6LC-0056626:/service000b/char000c]# notify on
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Notifying: yes
    Notify started
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Value:
    02 4f 03 0e 93                                   .O...
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Value:
    02 50 03 95 97                                   .P...
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Value:
    02 51 03 2c 9c                                   .Q.,.
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Value:
    02 53 03 d4 a1                                   .S...
    [CHG] Attribute /org/bluez/hci0/dev_XX_XX_XX_XX_XX_XX/service000b/char000c Value:
    02 54 03 a8 a4                                   .T...

