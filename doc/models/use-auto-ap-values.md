
# Use Auto Ap Values

## Structure

`UseAutoApValues`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | If accept is true, accepts placement for devices in list otherwise. If false, reject for devices in list.<br>**Default**: `false` |
| `For` | [`*models.UseAutoApValuesForEnum`](../../doc/models/use-auto-ap-values-for-enum.md) | Optional | The selector to choose auto placement or auto orientation.<br>**Default**: `"placement"` |
| `Macs` | `[]string` | Optional | A list of macs to accept/reject. If a list is not provided the API will accept/reject for the full map. |

## Example (as JSON)

```json
{
  "accept": false,
  "for": "placement",
  "macs": [
    "macs3",
    "macs2",
    "macs1"
  ]
}
```

