
# Optic Port Config Port

## Structure

`OpticPortConfigPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channelized` | `*bool` | Optional | Enable channelization<br><br>**Default**: `false` |
| `Speed` | `*string` | Optional | Interface speed (e.g. `25g`, `50g`), use the chassis speed by default |

## Example (as JSON)

```json
{
  "channelized": false,
  "speed": "25g"
}
```

