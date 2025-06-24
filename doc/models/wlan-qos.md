
# Wlan Qos

*This model accepts additional fields of type interface{}.*

## Structure

`WlanQos`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Class` | [`*models.WlanQosClassEnum`](../../doc/models/wlan-qos-class-enum.md) | Optional | enum: `background`, `best_effort`, `video`, `voice`<br><br>**Default**: `"best_effort"` |
| `Overwrite` | `*bool` | Optional | Whether to overwrite QoS<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "class": "best_effort",
  "overwrite": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

