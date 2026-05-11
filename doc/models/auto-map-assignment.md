
# Auto Map Assignment

## Structure

`AutoMapAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dryrun` | `*bool` | Optional | If `true`, validates the site's APs without starting the map assignment process. Returns device validity and estimated runtime.<br><br>**Default**: `false` |
| `ForceCollection` | `*bool` | Optional | If `true`, forces data collection via orchestration. If `false`, attempts to use existing BLE data first.<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "dryrun": false,
  "force_collection": false
}
```

