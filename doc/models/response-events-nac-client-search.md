
# Response Events Nac Client Search

## Structure

`ResponseEventsNacClientSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.EventNacClient`](../../doc/models/event-nac-client.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1513176951,
  "limit": 10,
  "start": 1512572151,
  "total": 1,
  "results": [
    {
      "ap": "ap8",
      "auth_type": "auth_type4",
      "bssid": "bssid0",
      "device_mac": "device_mac0",
      "dryrun_nacrule_id": "00000112-0000-0000-0000-000000000000"
    }
  ]
}
```

