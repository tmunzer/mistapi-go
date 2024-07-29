
# Wlan Qos

## Structure

`WlanQos`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Class` | [`*models.WlanQosClassEnum`](../../doc/models/wlan-qos-class-enum.md) | Optional | enum: `background`, `best_effort`, `video`, `voice`<br>**Default**: `"best_effort"` |
| `Overwrite` | `*bool` | Optional | whether to overwrite QoS<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "class": "best_effort",
  "overwrite": false
}
```

