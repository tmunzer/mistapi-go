
# Sso Mxedge Proxy Acct Server

*This model accepts additional fields of type interface{}.*

## Structure

`SsoMxedgeProxyAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `1813` |
| `Secret` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "port": 1813,
  "secret": "testing123",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

