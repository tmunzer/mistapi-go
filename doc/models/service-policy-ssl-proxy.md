
# Service Policy Ssl Proxy

For SRX-only

## Structure

`ServicePolicySslProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiphersCategory` | [`*models.SslProxyCiphersCategoryEnum`](../../doc/models/ssl-proxy-ciphers-category-enum.md) | Optional | enum: `medium`, `strong`, `weak`<br><br>**Default**: `"strong"` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "ciphers_category": "strong",
  "enabled": false
}
```

