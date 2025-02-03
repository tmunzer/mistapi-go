
# Stats Zone Details

Zone details statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsZoneDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Assets` | `[]string` | Optional | List of ble assets currently in the zone and when they entered |
| `ClientWaits` | [`models.StatsZoneDetailsClientWaits`](../../doc/models/stats-zone-details-client-waits.md) | Required | Client wait time right now |
| `Clients` | `[]string` | Optional | List of clients currently in the zone and when they entered |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `MapId` | `uuid.UUID` | Required | Map_id of the zone |
| `Name` | `string` | Required | Name of the zone |
| `NumClients` | `int` | Required | - |
| `NumSdkclients` | `int` | Required | SDK client wait time right now |
| `Sdkclients` | `[]string` | Optional | List of SDK Clients currently in the zone and when they entered |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_waits": {
    "avg": 1200,
    "max": 3610,
    "min": 600,
    "p95": 2800,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "Board Room",
  "num_clients": 80,
  "num_sdkclients": 0,
  "assets": [
    "assets2",
    "assets1",
    "assets0"
  ],
  "clients": [
    "clients2"
  ],
  "sdkclients": [
    "sdkclients0",
    "sdkclients1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

