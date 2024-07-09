
# Reponse Synthetictest Search

## Structure

`ReponseSynthetictestSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.SynthetictestInfo`](../../doc/models/synthetictest-info.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 220,
  "limit": 206,
  "results": [
    {
      "by": "user",
      "failed": false,
      "latency": 40,
      "port_id": "ge-0/0/2",
      "reason": "interface not ready to perform test",
      "rx_mbps": 322,
      "start_time": 1675718807,
      "timestamp": 1675718807,
      "tx_mbps": 199,
      "vlan_id": 20,
      "device_type": "switch",
      "mac": "mac0"
    }
  ],
  "start": 178,
  "total": 44,
  "next": "next0"
}
```
