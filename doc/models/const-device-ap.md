
# Const Device Ap

AP model capability definition returned by the constants API

## Structure

`ConstDeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApType` | `string` | Required | Internal AP platform type identifier |
| `Band24` | [`*models.ConstDeviceApBand24`](../../doc/models/const-device-ap-band-24.md) | Optional | 2.4 GHz radio capability limits for an AP model |
| `Band5` | [`*models.ConstDeviceApBand5`](../../doc/models/const-device-ap-band-5.md) | Optional | 5 GHz radio capability limits for an AP model |
| `Band6` | [`*models.ConstDeviceApBand6`](../../doc/models/const-device-ap-band-6.md) | Optional | 6 GHz radio capability limits for an AP model |
| `Band24Usages` | [`[]models.ConstDeviceApBand24UsageEnum`](../../doc/models/const-device-ap-band-24-usage-enum.md) | Optional | Supported 2.4 GHz radio usage modes for AP models |
| `CeDfsOk` | `*bool` | Optional | Whether DFS operation is allowed for this AP model under CE rules |
| `CiscoPace` | `*bool` | Optional | Whether the AP model supports Cisco PACE interoperability |
| `Description` | `*string` | Optional | Product description for the AP model |
| `DisallowedChannels` | `map[string][]int` | Optional | Property key is a list of country codes (e.g. "GB, DE") |
| `Display` | `*string` | Optional | User-facing model name shown for the AP |
| `Extio` | [`map[string]models.ConstDeviceApExtios`](../../doc/models/const-device-ap-extios.md) | Optional | Property key is the GPIO port name (e.g. "D0", "A1") |
| `FccDfsOk` | `*bool` | Optional | Whether DFS operation is allowed for this AP model under FCC rules |
| `Has11ax` | `*bool` | Optional | Whether the AP model supports 802.11ax |
| `HasCompass` | `*bool` | Optional | Whether the AP model includes a compass sensor |
| `HasExtAnt` | `*bool` | Optional | Whether the AP model supports external antennas |
| `HasExtio` | `*bool` | Optional | Whether the AP model exposes external I/O ports |
| `HasHeight` | `*bool` | Optional | Whether mounting height can be configured for this AP model |
| `HasModulePort` | `*bool` | Optional | Whether the AP model includes a module port |
| `HasPoeOut` | `*bool` | Optional | Whether the AP model supports PoE output |
| `HasScanningRadio` | `*bool` | Optional | Whether the AP model has scanning-radio capability |
| `HasSelectableRadio` | `*bool` | Optional | Whether the AP model has selectable radio modes |
| `HasUsb` | `*bool` | Optional | Whether the AP model includes a USB port |
| `HasVble` | `*bool` | Optional | Whether the AP model supports virtual BLE (vBLE) |
| `HasWifiBand24` | `*bool` | Optional | Whether the AP model supports 2.4 GHz Wi-Fi |
| `HasWifiBand5` | `*bool` | Optional | Whether the AP model supports 5 GHz Wi-Fi |
| `HasWifiBand6` | `*bool` | Optional | Whether the AP model supports 6 GHz Wi-Fi |
| `MaxPoeOut` | `*int` | Optional | Maximum PoE-out power budget supported by the AP model, in milliwatts |
| `MaxWlans` | `*int` | Optional | Maximum number of WLANs supported by this AP model |
| `Model` | `*string` | Optional | AP model identifier for this capability definition |
| `OtherDfsOk` | `*bool` | Optional | Whether DFS operation is allowed for this AP model in other regulatory domains |
| `Outdoor` | `*bool` | Optional | Whether the AP model is rated for outdoor deployment |
| `Radios` | `map[string]string` | Optional | Property key is the radio number (e.g. r0, r1, ...). Property value is the RF band (e.g. "24", "5", ...) |
| `SharedScanningRadio` | `*bool` | Optional | Whether scanning-radio capability shares a radio with client service |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `ap`<br><br>**Value**: `"ap"` |
| `Unmanaged` | `*bool` | Optional | Whether this AP model is listed as unmanaged in the constants catalog |
| `Vble` | [`*models.ConstDeviceApVble`](../../doc/models/const-device-ap-vble.md) | Optional | Virtual BLE (vBLE) capability settings for an AP model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceAp := models.ConstDeviceAp{
        ApType:               "jewel",
        Band24:               models.ToPointer(models.ConstDeviceApBand24{
            Band5ChannelsOp:      models.ToPointer("band5_channels_op4"),
            MaxClients:           models.ToPointer(176),
            MaxPower:             models.ToPointer(208),
            MinPower:             models.ToPointer(38),
        }),
        Band5:                models.ToPointer(models.ConstDeviceApBand5{
            MaxClients:           models.ToPointer(252),
            MaxPower:             models.ToPointer(132),
            MinPower:             models.ToPointer(142),
        }),
        Band6:                models.ToPointer(models.ConstDeviceApBand6{
            MaxClients:           models.ToPointer(154),
            MaxPower:             models.ToPointer(230),
            MinPower:             models.ToPointer(240),
        }),
        Band24Usages:         []models.ConstDeviceApBand24UsageEnum{
            models.ConstDeviceApBand24UsageEnum_ENUM24,
            models.ConstDeviceApBand24UsageEnum_ENUM5,
            models.ConstDeviceApBand24UsageEnum_ENUM6,
        },
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
    }

}
```

