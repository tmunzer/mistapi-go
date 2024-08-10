
# Auto Placement

## Structure

`AutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForceCollection` | `*bool` | Optional | * If `force_collection`==`false`: the API Iattempts to start localization with existing data.<br>* If `force_collection`==`true`: maintenance the API attempts to start orchestration.<br>**Default**: `false` |
| `Macs` | `[]string` | Optional | list of device macs |

## Example (as JSON)

```json
{
  "force_collection": false,
  "macs": [
    "macs9",
    "macs8",
    "macs7"
  ]
}
```

