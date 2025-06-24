
# Auto Orient

*This model accepts additional fields of type interface{}.*

## Structure

`AutoOrient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dryrun` | `*bool` | Optional | Set to `true` to perform an invalid AP check and provide an estimated run time without enqueuing the run into the auto orient service. |
| `ForceCollection` | `*bool` | Optional | If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.<br>If `force_collection`==`true`, the API attempts to start BLE orchestration.<br><br>**Default**: `false` |
| `Macs` | `[]string` | Optional | List of device macs |
| `Override` | `*bool` | Optional | Set to `true` to run auto orient even if there are invalid APs in the selected APs. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "force_collection": false,
  "dryrun": false,
  "macs": [
    "macs3",
    "macs4",
    "macs5"
  ],
  "override": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

