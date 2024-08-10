
# Utils Show Evpn Database

## Structure

`UtilsShowEvpnDatabase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | rate at which output will refresh<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 10` |
| `Mac` | `*string` | Optional | client mac filter |
| `PortId` | `*string` | Optional | interface name |

## Example (as JSON)

```json
{
  "duration": 0,
  "interval": 0,
  "mac": "f8c1165c6400",
  "port_id": "ge-0/0/0.0"
}
```

