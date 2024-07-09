
# Switch Stats If Stat

## Structure

`SwitchStatsIfStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ips` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `PortId` | `*string` | Optional | - |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/0",
  "ips": [
    "ips4",
    "ips5",
    "ips6"
  ],
  "rx_bytes": 58,
  "rx_pkts": 116,
  "tx_bytes": 196
}
```

