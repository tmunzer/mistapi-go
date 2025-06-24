
# Ap Radio Mac

*This model accepts additional fields of type interface{}.*

## Structure

`ApRadioMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `RadioMacs` | `[]string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b350001a0",
  "radio_macs": [
    "5c5b350001a0",
    "5c5b350001a1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

