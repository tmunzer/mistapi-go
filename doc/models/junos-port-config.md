
# Junos Port Config

Switch port config

## Structure

`JunosPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | To disable LACP support for the AE interface |
| `AeIdx` | `*int` | Optional | Users could force to use the designated AE name |
| `AeLacpSlow` | `*bool` | Optional | to use fast timeout<br>**Default**: `true` |
| `Aggregated` | `*bool` | Optional | **Default**: `false` |
| `Critical` | `*bool` | Optional | if want to generate port up/down alarm |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | if `speed` and `duplex` are specified, whether to disable autonegotiation<br>**Default**: `false` |
| `Duplex` | [`*models.JunosPortConfigDuplexEnum`](../../doc/models/junos-port-config-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br>**Default**: `"auto"` |
| `DynamicUsage` | `models.Optional[string]` | Optional | Enable dynamic usage for this port. Set to `dynamic` to enable. |
| `Esilag` | `*bool` | Optional | - |
| `Mtu` | `*int` | Optional | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation<br>**Default**: `1514` |
| `NoLocalOverwrite` | `*bool` | Optional | prevent helpdesk to override the port config |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br>**Default**: `"auto"` |
| `Usage` | `string` | Required | port usage name. If EVPN is used, use `evpn_uplink`or `evpn_downlink` |

## Example (as JSON)

```json
{
  "ae_lacp_slow": true,
  "aggregated": false,
  "disable_autoneg": false,
  "duplex": "auto",
  "mtu": 1514,
  "poe_disabled": false,
  "speed": "auto",
  "usage": "usage6",
  "ae_disable_lacp": false,
  "ae_idx": 244,
  "critical": false
}
```

