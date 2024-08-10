
# Response Guest Search

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

## Example (as JSON)

```json
{
  "end": 166,
  "limit": 4,
  "next": "next0",
  "results": [
    {
      "authorized_expiring_time": 1480704955.0,
      "authorized_time": 1480704355.0,
      "company": "abc",
      "email": "john@abc.com",
      "name": "John Smith",
      "ssid": "Guest-SSID",
      "wlan_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "authorized": false
    }
  ],
  "start": 124,
  "total": 98
}
```

