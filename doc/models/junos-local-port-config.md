
# Junos Local Port Config

Switch port config

## Structure

`JunosLocalPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Critical` | `*bool` | Optional | if want to generate port up/down alarm |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | if `speed` and `duplex` are specified, whether to disable autonegotiation<br>**Default**: `false` |
| `Duplex` | [`*models.JunosPortConfigDuplexEnum`](../../doc/models/junos-port-config-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br>**Default**: `"auto"` |
| `Mtu` | `*int` | Optional | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation<br>**Default**: `1514` |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `auto`<br>**Default**: `"auto"` |
| `Usage` | `string` | Required | port usage name.<br><br>If EVPN is used, use `evpn_uplink`or `evpn_downlink` |

## Example (as JSON)

```json
{
  "disable_autoneg": false,
  "duplex": "auto",
  "mtu": 1514,
  "poe_disabled": false,
  "speed": "auto",
  "usage": "usage8",
  "critical": false,
  "description": "description0"
}
```

