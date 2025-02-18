
# Response Mxtunnels Preempt Aps

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMxtunnelsPreemptAps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreemptedAps` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "preempted_aps": [
    "preempted_aps8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

