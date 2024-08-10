
# Const Device Model

## Structure

`ConstDeviceModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApType` | `*string` | Optional | - |
| `Band24` | [`*models.ConstDeviceApBand24`](../../doc/models/const-device-ap-band-24.md) | Optional | - |
| `Band5` | [`*models.ConstDeviceApBand5`](../../doc/models/const-device-ap-band-5.md) | Optional | - |
| `Band6` | [`*models.ConstDeviceApBand5`](../../doc/models/const-device-ap-band-5.md) | Optional | - |
| `Band24Usages` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
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
| `Type` | `*string` | Optional | **Default**: `"ap"` |
| `Unmanaged` | `*bool` | Optional | - |
| `Vble` | [`*models.ConstDeviceApVble`](../../doc/models/const-device-ap-vble.md) | Optional | - |
| `Alias` | `*string` | Optional | - |
| `Defaults` | [`*models.ConstDeviceSwitchDefault1`](../../doc/models/const-device-switch-default-1.md) | Optional | Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...) |
| `EvolvedOs` | `*bool` | Optional | **Default**: `false` |
| `EvpnRiType` | `*string` | Optional | - |
| `Experimental` | `*bool` | Optional | **Default**: `false` |
| `FansPluggable` | `*bool` | Optional | **Default**: `false` |
| `HasBgp` | `*bool` | Optional | **Default**: `false` |
| `HasEts` | `*bool` | Optional | **Default**: `false` |
| `HasEvpn` | `*bool` | Optional | **Default**: `false` |
| `HasIrb` | `*bool` | Optional | **Default**: `false` |
| `HasSnapshot` | `*bool` | Optional | **Default**: `true` |
| `HasVc` | `*bool` | Optional | **Default**: `true` |
| `Modular` | `*bool` | Optional | **Default**: `false` |
| `NoShapingRate` | `*bool` | Optional | **Default**: `false` |
| `NumberFans` | `*int` | Optional | - |
| `OcDevice` | `*bool` | Optional | **Default**: `false` |
| `OobInterface` | `*string` | Optional | - |
| `PacketActionDropOnly` | `*bool` | Optional | **Default**: `false` |
| `Pic` | `map[string]string` | Optional | Object Key is the PIC number |
| `SubRequired` | `*string` | Optional | - |
| `HaNode0Fpc` | `*int` | Optional | - |
| `HaNode1Fpc` | `*int` | Optional | - |
| `HasFxp0` | `*bool` | Optional | **Default**: `true` |
| `HasHaControl` | `*bool` | Optional | **Default**: `false` |
| `HasHaData` | `*bool` | Optional | **Default**: `false` |
| `IrbDisabledByDefault` | `*bool` | Optional | **Default**: `false` |
| `Ports` | [`*models.ConstDeviceGatewayPorts`](../../doc/models/const-device-gateway-ports.md) | Optional | Object Key is the interface name (e.g. "ge-0/0/1", ...) |
| `T128Device` | `*bool` | Optional | **Default**: `false` |

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
  "alias": "EX4100-48P-CHAS",
  "evolved_os": false,
  "evpn_ri_type": "mac-vrf",
  "experimental": false,
  "fans_pluggable": true,
  "has_bgp": true,
  "has_ets": false,
  "has_evpn": true,
  "has_irb": true,
  "has_snapshot": true,
  "has_vc": true,
  "modular": false,
  "no_shaping_rate": false,
  "number_fans": 2,
  "oc_device": true,
  "oob_interface": "re0:mgmt-0, re1:mgmt-0",
  "packet_action_drop_only": false,
  "pic": {
    "0": "ge*48",
    "1": "qsfp+*4",
    "2": "sfp+*4 (uplink)"
  },
  "has_fxp0": true,
  "has_ha_control": false,
  "has_ha_data": false,
  "irb_disabled_by_default": false,
  "t128_device": false,
  "band24": {
    "band5_channels_op": "band5_channels_op4",
    "max_clients": 176,
    "max_power": 208,
    "min_power": 38
  },
  "band5": {
    "max_clients": 252,
    "max_power": 132,
    "min_power": 142
  },
  "band6": {
    "max_clients": 154,
    "max_power": 230,
    "min_power": 240
  },
  "band_24_usages": "6"
}
```

