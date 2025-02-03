
# Utils Bounce Port

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsBouncePort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of ports to bounce |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": [
    "ports1",
    "ports2"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

