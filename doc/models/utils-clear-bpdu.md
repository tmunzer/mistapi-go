
# Utils Clear Bpdu

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearBpdu`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of ports on which to clear the detected BPDU error, or `all` for all ports |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": [
    "ports3",
    "ports2",
    "ports1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

