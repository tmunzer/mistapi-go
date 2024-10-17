
# Org Service Policy Ssl Proxy

for SRX-only

## Structure

`OrgServicePolicySslProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiphersCatagory` | [`*models.SslProxyCiphersCatagoryEnum`](../../doc/models/ssl-proxy-ciphers-catagory-enum.md) | Optional | enum: `medium`, `strong`, `weak`<br>**Default**: `"strong"` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "ciphers_catagory": "strong",
  "enabled": false
}
```

