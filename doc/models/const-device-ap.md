
# Const Device Ap

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApType` | `string` | Required | - |
| `Band24` | [`*models.ConstDeviceApBand24`](../../doc/models/const-device-ap-band-24.md) | Optional | - |
| `Band5` | [`*models.ConstDeviceApBand5`](../../doc/models/const-device-ap-band-5.md) | Optional | - |
| `Band6` | [`*models.ConstDeviceApBand5`](../../doc/models/const-device-ap-band-5.md) | Optional | - |
| `Band24Usages` | [`[]models.ConstDeviceApBand24UsageEnum`](../../doc/models/const-device-ap-band-24-usage-enum.md) | Optional | - |
| `CeDfsOk` | `*bool` | Optional | - |
| `CiscoPace` | `*bool` | Optional | - |
| `Description` | `*string` | Optional | - |
| `DisallowedChannels` | `map[string]interface{}` | Optional | Property key is a list of country codes (e.g. "GB, DE") |
| `Display` | `*string` | Optional | - |
| `Extio` | [`map[string]models.ConstDeviceApExtios`](../../doc/models/const-device-ap-extios.md) | Optional | Property key is the GPIO port name (e.g. "D0", "A1") |
| `FccDfsOk` | `*bool` | Optional | - |
| `Has11ax` | `*bool` | Optional | - |
| `HasCompass` | `*bool` | Optional | - |
| `HasExtAnt` | `*bool` | Optional | - |
| `HasExtio` | `*bool` | Optional | - |
| `HasHeight` | `*bool` | Optional | - |
| `HasModulePort` | `*bool` | Optional | - |
| `HasPoeOut` | `*bool` | Optional | - |
| `HasScanningRadio` | `*bool` | Optional | - |
| `HasSelectableRadio` | `*bool` | Optional | - |
| `HasUsb` | `*bool` | Optional | - |
| `HasVble` | `*bool` | Optional | - |
| `HasWifiBand24` | `*bool` | Optional | - |
| `HasWifiBand5` | `*bool` | Optional | - |
| `HasWifiBand6` | `*bool` | Optional | - |
| `MaxPoeOut` | `*int` | Optional | - |
| `MaxWlans` | `*int` | Optional | - |
| `Model` | `*string` | Optional | - |
| `OtherDfsOk` | `*bool` | Optional | - |
| `Outdoor` | `*bool` | Optional | - |
| `Radios` | `map[string]string` | Optional | Property key is the radio number (e.g. r0, r1, ...). Property value is the RF band (e.g. "24", "5", ...) |
| `SharedScanningRadio` | `*bool` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `ap`<br>**Default**: `"ap"` |
| `Unmanaged` | `*bool` | Optional | - |
| `Vble` | [`*models.ConstDeviceApVble`](../../doc/models/const-device-ap-vble.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_type": "jewel",
  "ce_dfs_ok": true,
  "description": "AP-45",
  "display": "AP45",
  "fcc_dfs_ok": true,
  "has_compass": false,
  "has_extio": false,
  "has_height": false,
  "has_poe_out": true,
  "has_scanning_radio": true,
  "has_selectable_radio": true,
  "has_vble": true,
  "has_wifi_band24": true,
  "has_wifi_band5": true,
  "has_wifi_band6": true,
  "max_poe_out": 15400,
  "model": "AP45",
  "other_dfs_ok": true,
  "radios": {
    "r0": "6",
    "r1": "5",
    "r2": "24"
  },
  "type": "ap",
  "band24": {
    "band5_channels_op": "band5_channels_op4",
    "max_clients": 176,
    "max_power": 208,
    "min_power": 38,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band5": {
    "max_clients": 252,
    "max_power": 132,
    "min_power": 142,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band6": {
    "max_clients": 154,
    "max_power": 230,
    "min_power": 240,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_24_usages": [
    "6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

