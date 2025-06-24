
# Ap Client Bridge Auth

*This model accepts additional fields of type interface{}.*

## Structure

`ApClientBridgeAuth`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Psk` | `*string` | Optional | **Constraints**: *Minimum Length*: `8`, *Maximum Length*: `63` |
| `Type` | [`*models.ApClientBridgeAuthTypeEnum`](../../doc/models/ap-client-bridge-auth-type-enum.md) | Optional | wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`<br><br>**Default**: `"psk"`<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "psk": "foryoureyesonly",
  "type": "psk",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

