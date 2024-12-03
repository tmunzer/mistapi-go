
# Insight Rogue Client

*This model accepts additional fields of type interface{}.*

## Structure

`InsightRogueClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Annotation` | `string` | Required | - |
| `ApMac` | `string` | Required | - |
| `AvgRssi` | `float64` | Required | - |
| `Band` | `string` | Required | - |
| `Bssid` | `string` | Required | - |
| `ClientMac` | `string` | Required | - |
| `NumAps` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "annotation": "annotation4",
  "ap_mac": "ap_mac2",
  "avg_rssi": 56.1,
  "band": "band2",
  "bssid": "bssid4",
  "client_mac": "client_mac8",
  "num_aps": 230,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

