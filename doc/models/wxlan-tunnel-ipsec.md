
# Wxlan Tunnel Ipsec

IPSec-related configurations; requires DMVPN be enabled

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanTunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether ipsec is enabled, requires DMVPN be enabled<br><br>**Default**: `false` |
| `Psk` | `string` | Required | IPSec pre-shared key |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "psk": "psk8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

