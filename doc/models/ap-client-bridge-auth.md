
# Ap Client Bridge Auth

## Structure

`ApClientBridgeAuth`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Psk` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Type` | [`*models.ApClientBridgeAuthTypeEnum`](../../doc/models/ap-client-bridge-auth-type-enum.md) | Optional | wpa2-AES/CCMPp is assumed when `type`==`psk`<br>**Default**: `"psk"`<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "psk": "foryoureyesonly",
  "type": "psk"
}
```
