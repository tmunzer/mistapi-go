
# Response Iot Endpoints Search

## Structure

`ResponseIotEndpointsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Results` | [`[]models.IotendpointStats`](../../doc/models/iotendpoint-stats.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 1.74,
  "results": [
    {
      "ap_mac": "5c5b350e0001",
      "id": "63f9e299182b63f9",
      "mac": "63f9e299182b63f9",
      "mfg": "Assa Abloy",
      "model": "Assa Abloy",
      "type": "zigbee",
      "lqi": 168
    }
  ],
  "start": 213.8,
  "total": 90
}
```

