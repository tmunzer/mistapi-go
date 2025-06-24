
# Auto Placement

*This model accepts additional fields of type interface{}.*

## Structure

`AutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dryrun` | `*bool` | Optional | Set to `true` to perform an invalid AP check and provide an estimated run time without enqueuing the run into the auto placement service.<br><br>**Default**: `false` |
| `ForceCollection` | `*bool` | Optional | * If `force_collection`==`false`: the API attempts to start localization with existing data.<br>* If `force_collection`==`true`: maintenance the API attempts to start orchestration.<br><br>**Default**: `false` |
| `Macs` | `[]string` | Optional | List of device macs |
| `Override` | `*bool` | Optional | Set to `true` to run auto placement even if there are invalid APs in the selected APs.<br><br>**Default**: `false` |
| `UwbOnly` | `*bool` | Optional | If set to `true`, the service shall be using UWB ranging without invoking Maintenance Mode.<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dryrun": false,
  "force_collection": false,
  "override": false,
  "uwb_only": false,
  "macs": [
    "macs9",
    "macs8",
    "macs7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

