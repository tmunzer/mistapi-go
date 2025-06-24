
# Service Policy Ssl Proxy

For SRX-only

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicySslProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiphersCategory` | [`*models.SslProxyCiphersCategoryEnum`](../../doc/models/ssl-proxy-ciphers-category-enum.md) | Optional | enum: `medium`, `strong`, `weak`<br><br>**Default**: `"strong"` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ciphers_category": "strong",
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

