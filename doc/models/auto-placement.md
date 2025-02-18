
# Auto Placement

*This model accepts additional fields of type interface{}.*

## Structure

`AutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForceCollection` | `*bool` | Optional | * If `force_collection`==`false`: the API Iattempts to start localization with existing data.<br>* If `force_collection`==`true`: maintenance the API attempts to start orchestration.<br>**Default**: `false` |
| `Macs` | `[]string` | Optional | List of device macs |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "force_collection": false,
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

