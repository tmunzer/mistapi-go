
# Wxlan Tunnel Ipsec

IPSec-related configurations; requires DMVPN be enabled

## Structure

`WxlanTunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether ipsec is enabled, requires DMVPN be enabled<br>**Default**: `false` |
| `Psk` | `string` | Required | ipsec pre-shared key |

## Example (as JSON)

```json
{
  "enabled": false,
  "psk": "psk4"
}
```

