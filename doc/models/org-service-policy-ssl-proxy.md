
# Org Service Policy Ssl Proxy

For SRX-only

*This model accepts additional fields of type interface{}.*

## Structure

`OrgServicePolicySslProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiphersCatagory` | [`*models.SslProxyCiphersCatagoryEnum`](../../doc/models/ssl-proxy-ciphers-catagory-enum.md) | Optional | enum: `medium`, `strong`, `weak`<br>**Default**: `"strong"` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ciphers_catagory": "strong",
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

