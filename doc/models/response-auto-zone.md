
# Response Auto Zone

## Structure

`ResponseAutoZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`*models.ResponseAutoZoneStatusEnum`](../../doc/models/response-auto-zone-status-enum.md) | Optional | The status for the auto zones service for a given map. enum:<br><br>* not_started: The auto zones service has not been run on this map or the results were cleared by the user<br>* in_progress: The auto zones service is currently in progress<br>* awaiting_review: The auto zones service has completed and suggested location zones to be added to the map<br>* error: There was an error with the auto zones service |
| `Zones` | [`[]models.ResponseAutoZoneZone`](../../doc/models/response-auto-zone-zone.md) | Optional | - |

## Example (as JSON)

```json
{
  "status": "not_started",
  "zones": [
    {
      "name": "name8",
      "vertices": [
        {
          "x": 16,
          "y": 88
        },
        {
          "x": 16,
          "y": 88
        }
      ]
    },
    {
      "name": "name8",
      "vertices": [
        {
          "x": 16,
          "y": 88
        },
        {
          "x": 16,
          "y": 88
        }
      ]
    },
    {
      "name": "name8",
      "vertices": [
        {
          "x": 16,
          "y": 88
        },
        {
          "x": 16,
          "y": 88
        }
      ]
    }
  ]
}
```

