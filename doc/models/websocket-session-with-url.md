
# Websocket Session With Url

*This model accepts additional fields of type interface{}.*

## Structure

`WebsocketSessionWithUrl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Session` | `string` | Required | - |
| `Url` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "session": "19e73828-937f-05e6-f709-e29efdb0a82b",
  "url": "wss://api-ws.mist.com/ssh?jwt=xxxx",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

