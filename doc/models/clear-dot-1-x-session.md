
# Clear Dot 1 X Session

*This model accepts additional fields of type interface{}.*

## Structure

`ClearDot1xSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of port IDs where the dot1x session must be cleared. Use `all` to clear sessions on all ports. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": [
    "ge-0/0/0",
    "ge-0/0/1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

