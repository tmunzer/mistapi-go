
# Ospf Peer Stats Search Results Items

*This model accepts additional fields of type interface{}.*

## Structure

`OspfPeerStatsSearchResultsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeadTime` | `*int` | Optional | Activity timer |
| `Mac` | `*string` | Optional | Router MAC address |
| `OrgId` | `*uuid.UUID` | Optional | Router org ID |
| `PeerIp` | `*string` | Optional | Neighbor address (IP) |
| `PortId` | `*string` | Optional | Interface name |
| `Priority` | `*int` | Optional | Neighbor priority, 0-255<br><br>**Constraints**: `>= 0`, `<= 255` |
| `SiteId` | `*uuid.UUID` | Optional | Router site ID |
| `State` | `*string` | Optional | Eg. full, down, 2way, init, exstart, exchange, loading |
| `Timestamp` | `*float64` | Optional | Sampling time (in epoch seconds) |
| `Up` | `*bool` | Optional | True if state is full |
| `VrfName` | `*string` | Optional | Instance name, e.g. master |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "dead_time": 194,
  "mac": "mac4",
  "peer_ip": "peer_ip0",
  "port_id": "port_id0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

