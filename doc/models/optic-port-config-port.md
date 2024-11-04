
# Optic Port Config Port

## Structure

`OpticPortConfigPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channelized` | `*bool` | Optional | enable channelization<br>**Default**: `false` |
| `Speed` | `*string` | Optional | interface speed (e.g. `25g`, `50g`), use the chassis speed by default |

## Example (as JSON)

```json
{
  "channelized": false,
  "speed": "25g"
}
```

