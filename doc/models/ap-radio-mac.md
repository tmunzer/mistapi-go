
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
  "mac": "mac4",
  "radio_macs": [
    "radio_macs7",
    "radio_macs8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

