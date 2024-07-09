
# Ap Radio Mac

## Structure

`ApRadioMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `RadioMacs` | `[]interface{}` | Required | - |

## Example (as JSON)

```json
{
  "mac": "mac8",
  "radio_macs": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ]
}
```

