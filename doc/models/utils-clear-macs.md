
# Utils Clear Macs

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearMacs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of ports on which to clear mac addresses. must include logical unit number |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": [
    "ports9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

