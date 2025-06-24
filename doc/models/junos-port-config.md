
# Junos Port Config

Switch port config

*This model accepts additional fields of type interface{}.*

## Structure

`JunosPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | To disable LACP support for the AE interface |
| `AeIdx` | `*int` | Optional | Users could force to use the designated AE name |
| `AeLacpSlow` | `*bool` | Optional | To use fast timeout |
| `Aggregated` | `*bool` | Optional | **Default**: `false` |
| `Critical` | `*bool` | Optional | To generate port up/down alarm<br><br>**Default**: `false` |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | If `speed` and `duplex` are specified, whether to disable autonegotiation<br><br>**Default**: `false` |
| `Duplex` | [`*models.JunosPortConfigDuplexEnum`](../../doc/models/junos-port-config-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `DynamicUsage` | `models.Optional[string]` | Optional | Enable dynamic usage for this port. Set to `dynamic` to enable. |
| `Esilag` | `*bool` | Optional | - |
| `Mtu` | `*int` | Optional | Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation<br><br>**Default**: `1514` |
| `NoLocalOverwrite` | `*bool` | Optional | Prevent helpdesk to override the port config<br><br>**Default**: `true` |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `PortNetwork` | `*string` | Optional | Required if `usage`==`vlan_tunnel`. Q-in-Q tunneling using All-in-one bundling. This also enables standard L2PT for interfaces that are not encapsulation tunnel interfaces and uses MAC rewrite operation. [View more information](https://www.juniper.net/documentation/us/en/software/junos/multicast-l2/topics/topic-map/q-in-q.html#id-understanding-qinq-tunneling-and-vlan-translation) |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `Usage` | `string` | Required | Port usage name. For Q-in-Q, use `vlan_tunnel`. If EVPN is used, use `evpn_uplink`or `evpn_downlink` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aggregated": false,
  "critical": false,
  "disable_autoneg": false,
  "duplex": "auto",
  "mtu": 1514,
  "no_local_overwrite": true,
  "poe_disabled": false,
  "speed": "auto",
  "usage": "usage6",
  "ae_disable_lacp": false,
  "ae_idx": 244,
  "ae_lacp_slow": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

