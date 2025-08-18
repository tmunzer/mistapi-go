
# Vrrp Config Group

*This model accepts additional fields of type interface{}.*

## Structure

`VrrpConfigGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Preempt` | `*bool` | Optional | If `true`, allow preemption (a backup router can preempt a primary router)<br><br>**Default**: `false` |
| `Priority` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "preempt": false,
  "priority": 24,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

