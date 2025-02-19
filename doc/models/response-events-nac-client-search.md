
# Response Events Nac Client Search

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "auth_type": "peap-tls",
      "bssid": "bssid0",
      "device_mac": "device_mac0",
      "dryrun_nacrule_id": "00000112-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap": "ap8",
      "auth_type": "peap-tls",
      "bssid": "bssid0",
      "device_mac": "device_mac0",
      "dryrun_nacrule_id": "00000112-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

