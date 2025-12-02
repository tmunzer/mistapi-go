
# Response Events Nac Client Search

## Structure

`ResponseEventsNacClientSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.NacClientEvent`](../../doc/models/nac-client-event.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1513176951,
  "limit": 10,
  "start": 1512572151,
  "total": 1,
  "next": "next4",
  "results": [
    {
      "ap": "ap8",
      "auth_type": "eap-teap",
      "bssid": "bssid0",
      "client_type": "wired",
      "device_mac": "device_mac0"
    },
    {
      "ap": "ap8",
      "auth_type": "eap-teap",
      "bssid": "bssid0",
      "client_type": "wired",
      "device_mac": "device_mac0"
    }
  ]
}
```

