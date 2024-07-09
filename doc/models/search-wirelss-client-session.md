
# Search Wirelss Client Session

## Structure

`SearchWirelssClientSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.WirelssClientSession`](../../doc/models/wirelss-client-session.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 104.86,
  "limit": 76,
  "results": [
    {
      "ap": "ap8",
      "band": "band8",
      "client_manufacture": "client_manufacture0",
      "connect": 198,
      "disconnect": 148,
      "duration": 104.42,
      "mac": "mac0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ssid": "ssid6",
      "timestamp": 2.64,
      "wlan_id": "00000742-0000-0000-0000-000000000000",
      "for_site": false,
      "tags": [
        "tags1",
        "tags2"
      ]
    }
  ],
  "start": 60.92,
  "total": 238,
  "next": "next6"
}
```

