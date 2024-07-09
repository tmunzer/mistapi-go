
# Utils Show Arp

## Structure

`UtilsShowArp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | rate at which output will refresh<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 10` |
| `Ip` | `*string` | Optional | IP Address |
| `PortId` | `*string` | Optional | device Port ID |
| `Vrf` | `*string` | Optional | VRF Name |

## Example (as JSON)

```json
{
  "duration": 0,
  "interval": 0,
  "ip": "192.168.30.7",
  "port_id": "ge-0/0/0.0",
  "vrf": "guest"
}
```

