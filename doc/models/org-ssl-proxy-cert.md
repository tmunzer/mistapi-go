
# Org Ssl Proxy Cert

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSslProxyCert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cert": "-----BEGIN CERTIFICATE-----\\nMIIowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

