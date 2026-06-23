
# Const Device Model

Device model definition returned by the constants API

## Class Name

`ConstDeviceModel`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ConstDeviceAp`](../../../doc/models/const-device-ap.md) | models.ConstDeviceModelContainer.FromConstDeviceAp(models.ConstDeviceAp constDeviceAp) |
| [`models.ConstDeviceSwitch`](../../../doc/models/const-device-switch.md) | models.ConstDeviceModelContainer.FromConstDeviceSwitch(models.ConstDeviceSwitch constDeviceSwitch) |
| [`models.ConstDeviceGateway`](../../../doc/models/const-device-gateway.md) | models.ConstDeviceModelContainer.FromConstDeviceGateway(models.ConstDeviceGateway constDeviceGateway) |

## models.ConstDeviceAp

### Initialization Code

#### Example

```go
value := models.ConstDeviceModelContainer.FromConstDeviceAp(models.ConstDeviceAp{
    ApType:               "jewel",
    CeDfsOk:              models.ToPointer(true),
    Description:          models.ToPointer("AP-45"),
    Display:              models.ToPointer("AP45"),
    FccDfsOk:             models.ToPointer(true),
    HasCompass:           models.ToPointer(false),
    HasExtio:             models.ToPointer(false),
    HasHeight:            models.ToPointer(false),
    HasPoeOut:            models.ToPointer(true),
    HasScanningRadio:     models.ToPointer(true),
    HasSelectableRadio:   models.ToPointer(true),
    HasVble:              models.ToPointer(true),
    HasWifiBand24:        models.ToPointer(true),
    HasWifiBand5:         models.ToPointer(true),
    HasWifiBand6:         models.ToPointer(true),
    MaxPoeOut:            models.ToPointer(15400),
    Model:                models.ToPointer("AP45"),
    OtherDfsOk:           models.ToPointer(true),
    Radios:               map[string]string{
        "r0": "6",
        "r1": "5",
        "r2": "24",
    },
    Type:                 "ap",
})
```

## models.ConstDeviceSwitch

### Initialization Code

#### Example

```go
value := models.ConstDeviceModelContainer.FromConstDeviceSwitch(models.ConstDeviceSwitch{
    Alias:                models.ToPointer("EX4100-48P-CHAS"),
    Description:          models.ToPointer("Juniper EX4100 Series"),
    Display:              models.ToPointer("EX4100-48P"),
    EvolvedOs:            models.ToPointer(false),
    EvpnRiType:           models.ToPointer("mac-vrf"),
    Experimental:         models.ToPointer(false),
    FansPluggable:        models.ToPointer(true),
    HasBgp:               models.ToPointer(true),
    HasEts:               models.ToPointer(false),
    HasEvpn:              models.ToPointer(true),
    HasIrb:               models.ToPointer(true),
    HasPoeOut:            models.ToPointer(true),
    HasSnapshot:          models.ToPointer(true),
    HasVc:                models.ToPointer(true),
    Model:                models.ToPointer("EX4100-48P"),
    Modular:              models.ToPointer(false),
    NoShapingRate:        models.ToPointer(false),
    NumberFans:           models.ToPointer(2),
    OcDevice:             models.ToPointer(true),
    OobInterface:         models.ToPointer("re0:mgmt-0, re1:mgmt-0"),
    PacketActionDropOnly: models.ToPointer(false),
    Pic:                  map[string]string{
        "0": "ge*48",
        "1": "qsfp+*4",
        "2": "sfp+*4 (uplink)",
    },
    Type:                 "switch",
})
```

## models.ConstDeviceGateway

### Initialization Code

#### Example

```go
value := models.ConstDeviceModelContainer.FromConstDeviceGateway(models.ConstDeviceGateway{
    Experimental:         models.ToPointer(false),
    FansPluggable:        models.ToPointer(true),
    HasBgp:               models.ToPointer(false),
    HasFxp0:              models.ToPointer(true),
    HasHaControl:         models.ToPointer(false),
    HasHaData:            models.ToPointer(false),
    HasIrb:               models.ToPointer(false),
    HasPoeOut:            models.ToPointer(true),
    HasSnapshot:          models.ToPointer(true),
    IrbDisabledByDefault: models.ToPointer(false),
    OcDevice:             models.ToPointer(false),
    T128Device:           models.ToPointer(false),
    Type:                 "gateway",
})
```

