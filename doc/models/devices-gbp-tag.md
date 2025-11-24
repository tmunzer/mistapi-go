
# Devices Gbp Tag

*This model accepts additional fields of type interface{}.*

## Structure

`DevicesGbpTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GbpTag` | `int` | Required | - |
| `Macs` | `[]string` | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "gbp_tag": 4,
  "macs": [
    "683b679ac024"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

