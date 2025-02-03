
# Stats Wired Client

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthState` | `*string` | Optional | Client authorization status<br>**Constraints**: *Minimum Length*: `1` |
| `DeviceId` | `*string` | Optional | Device ID the client is connected to<br>**Constraints**: *Minimum Length*: `1` |
| `EthPort` | `*string` | Optional | Port on AP where the wired client is connected<br>**Constraints**: *Minimum Length*: `1` |
| `LastSeen` | `*float64` | Optional | Time when last Tx/Rx observed |
| `Mac` | `string` | Required | Client mac<br>**Constraints**: *Minimum Length*: `1` |
| `RxBytes` | `*float64` | Optional | Amount of traffic sent to client since client connects |
| `RxPkts` | `*float64` | Optional | Amount of traffic sent to client since client connects |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TxBytes` | `*float64` | Optional | amount of traffic received from client since client connects |
| `TxPkts` | `*float64` | Optional | Amount of traffic received from client since client connects |
| `Uptime` | `*float64` | Optional | How long, in seconds, has the client been connected |
| `VlanId` | `*float64` | Optional | VLAN id, could be empty |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac0",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "auth_state": "auth_state4",
  "device_id": "device_id2",
  "eth_port": "eth_port4",
  "last_seen": 111.26,
  "rx_bytes": 77.14,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

