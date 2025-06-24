
# Optic Port Config Port

*This model accepts additional fields of type interface{}.*

## Structure

`OpticPortConfigPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channelized` | `*bool` | Optional | Enable channelization<br><br>**Default**: `false` |
| `Speed` | `*string` | Optional | Interface speed (e.g. `25g`, `50g`), use the chassis speed by default |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channelized": false,
  "speed": "25g",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

