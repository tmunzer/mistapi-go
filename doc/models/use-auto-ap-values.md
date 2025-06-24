
# Use Auto Ap Values

*This model accepts additional fields of type interface{}.*

## Structure

`UseAutoApValues`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | If accept is true, accepts placement for devices in list otherwise. If false, reject for devices in list.<br><br>**Default**: `false` |
| `For` | [`*models.UseAutoApValuesForEnum`](../../doc/models/use-auto-ap-values-for-enum.md) | Optional | The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`<br><br>**Default**: `"placement"` |
| `Macs` | `[]string` | Optional | A list of macs to accept/reject. If a list is not provided the API will accept/reject for the full map. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accept": false,
  "for": "placement",
  "macs": [
    "macs1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

