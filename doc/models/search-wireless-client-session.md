
# Search Wireless Client Session

*This model accepts additional fields of type interface{}.*

## Structure

`SearchWirelessClientSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.WirelessClientSession`](../../doc/models/wireless-client-session.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 21.38,
  "limit": 176,
  "results": [
    {
      "ap": "ap8",
      "band": "band8",
      "connect": 198,
      "disconnect": 148,
      "duration": 104.42,
      "mac": "mac0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ssid": "ssid6",
      "timestamp": 2.64,
      "wlan_id": "00000742-0000-0000-0000-000000000000",
      "client_manufacture": "client_manufacture0",
      "for_site": false,
      "tags": [
        "tags1",
        "tags2"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 233.44,
  "total": 82,
  "next": "next8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

