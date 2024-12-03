
# Response Verify Token Success

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseVerifyTokenSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | - |
| `InviteNotApplied` | `*bool` | Optional | - |
| `MinLength` | `*int` | Optional | - |
| `ReturnTo` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "detail6",
  "invite_not_applied": false,
  "min_length": 206,
  "return_to": "return_to6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

