
# Stats Wired Client

## Structure

`StatsWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Ttl` | `*float64` | Optional | TTL of the validity of the stat |
| `AuthState` | `*string` | Optional | client authorization status<br>**Constraints**: *Minimum Length*: `1` |
| `DeviceId` | `*string` | Optional | Device ID the client is connected to<br>**Constraints**: *Minimum Length*: `1` |
| `EthPort` | `*string` | Optional | port on AP where the wired client is connected<br>**Constraints**: *Minimum Length*: `1` |
| `LastSeen` | `*float64` | Optional | time when last Tx/Rx observed |
| `Mac` | `string` | Required | client mac<br>**Constraints**: *Minimum Length*: `1` |
| `RxBytes` | `*float64` | Optional | amount of traffic sent to client since client connects |
| `RxPkts` | `*float64` | Optional | amount of traffic sent to client since client connects |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TxBytes` | `*float64` | Optional | amount of traffic received from client since client connects |
| `TxPkts` | `*float64` | Optional | amount of traffic received from client since client connects |
| `Uptime` | `*float64` | Optional | how long, in seconds, has the client been connected |
| `VlanId` | `*float64` | Optional | vlan id, could be empty |

## Example (as JSON)

```json
{
  "mac": "mac2",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "_id": "_id2",
  "_ttl": 236.08,
  "auth_state": "auth_state8",
  "device_id": "device_id4",
  "eth_port": "eth_port6"
}
```

