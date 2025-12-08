
# Wxlan Tunnel Ipsec

IPSec-related configurations; requires DMVPN be enabled

## Structure

`WxlanTunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether ipsec is enabled, requires DMVPN be enabled<br><br>**Default**: `false` |
| `Psk` | `string` | Required | IPSec pre-shared key |

## Example (as JSON)

```json
{
  "enabled": false,
  "psk": "psk8"
}
```

