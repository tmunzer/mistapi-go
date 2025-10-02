
# Clear Dot 1 X Session

*This model accepts additional fields of type interface{}.*

## Structure

`ClearDot1xSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortId` | `*string` | Optional | ID of the port where the dot1x session must be cleared. If not provided, the sessions on all the port will be cleared. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

