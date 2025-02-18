
# Response Guest Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseGuestSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `string` | Required | - |
| `Results` | [`[]models.Guest`](../../doc/models/guest.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 166,
  "limit": 4,
  "next": "next0",
  "results": [
    {
      "authorized": true,
      "authorized_expiring_time": 1480704955.0,
      "authorized_time": 1480704355,
      "company": "abc",
      "email": "john@abc.com",
      "minutes": 1440,
      "name": "John Smith",
      "ssid": "Guest-SSID",
      "wlan_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "access_code_email": "access_code_email8",
      "ap_mac": "ap_mac8",
      "auth_method": "auth_method0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 124,
  "total": 98,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

