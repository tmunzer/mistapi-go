
# Zone Stats Details

Zone details statistics

## Structure

`ZoneStatsDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Assets` | `[]string` | Optional | list of ble assets currently in the zone and when they entered |
| `ClientWaits` | [`models.ZoneStatsDetailsClientWaits`](../../doc/models/zone-stats-details-client-waits.md) | Required | client wait time right now |
| `Clients` | `[]string` | Optional | list of clients currently in the zone and when they entered |
| `Id` | `uuid.UUID` | Required | id of the zone |
| `MapId` | `uuid.UUID` | Required | map_id of the zone |
| `Name` | `string` | Required | name of the zone |
| `NumClients` | `int` | Required | - |
| `NumSdkclients` | `int` | Required | sdkclient wait time right now |
| `Sdkclients` | `[]string` | Optional | list of sdkclients currently in the zone and when they entered |

## Example (as JSON)

```json
{
  "client_waits": {
    "avg": 1200,
    "max": 3610,
    "min": 600,
    "p95": 2800
  },
  "id": "8ac84899-32db-6327-334c-9b6d58544cfe",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "Board Room",
  "num_clients": 80,
  "num_sdkclients": 0,
  "assets": [
    "assets0"
  ],
  "clients": [
    "clients0",
    "clients1",
    "clients2"
  ],
  "sdkclients": [
    "sdkclients6"
  ]
}
```

