
# Utils Release Dhcp Leases

Note:

* valid combinations for Junos:
  * `port_id`
  * `macs` + `network`
* valid combinations for SSR:
  * `port_id`
  * `macs` + `network`
  * `port_id` + `network`
  * `network`
* if network or port_id is specified and macs is empty, it means all clients under network or port_id

## Structure

`UtilsReleaseDhcpLeases`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `[]string` | Optional | A list of client macs to be released |
| `Network` | `*string` | Optional | The network for the leases IPs to be released |
| `PortId` | `string` | Required | The nework interface on which to release the current DHCP release<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "mac": [
    "90ec77aabbcc",
    "90ec77aabbdd"
  ],
  "network": "guest",
  "port_id": "ge-0/0/1.10"
}
```

