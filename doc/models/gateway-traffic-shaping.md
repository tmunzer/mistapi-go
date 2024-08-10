
# Gateway Traffic Shaping

## Structure

`GatewayTrafficShaping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClassPercentages` | `[]int` | Optional | percentages for differet class of traffic: high / medium / low / best-effort<br>sum must be equal to 100 |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "class_percentages": [
    201
  ]
}
```

