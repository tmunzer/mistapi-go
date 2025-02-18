
# Response Assign Success

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAssignSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "success": [
    "success4"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

