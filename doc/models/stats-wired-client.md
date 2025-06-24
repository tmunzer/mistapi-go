
# Stats Wired Client

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthState` | `*string` | Optional | Client authorization status<br><br>**Constraints**: *Minimum Length*: `1` |
| `DeviceId` | `*string` | Optional | Device ID the client is connected to<br><br>**Constraints**: *Minimum Length*: `1` |
| `EthPort` | `*string` | Optional | Port on AP where the wired client is connected<br><br>**Constraints**: *Minimum Length*: `1` |
| `LastSeen` | `*float64` | Optional | Time when last Tx/Rx observed |
| `Mac` | `string` | Required | Client mac<br><br>**Constraints**: *Minimum Length*: `1` |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Uptime` | `*float64` | Optional | How long, in seconds, has the client been connected |
| `VlanId` | `*float64` | Optional | VLAN id, could be empty |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac0",
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "auth_state": "auth_state4",
  "device_id": "device_id2",
  "eth_port": "eth_port4",
  "last_seen": 111.26,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

