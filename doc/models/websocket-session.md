
# Websocket Session

*This model accepts additional fields of type interface{}.*

## Structure

`WebsocketSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Session` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "session": "19e73828-937f-05e6-f709-e29efdb0a82b",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

