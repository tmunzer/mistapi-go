
# Tunnel Provider Options

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Jse` | [`*models.TunnelProviderOptionsJse`](../../doc/models/tunnel-provider-options-jse.md) | Optional | for jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added |
| `Zscaler` | [`*models.TunnelProviderOptionsZscaler`](../../doc/models/tunnel-provider-options-zscaler.md) | Optional | for zscaler-ipsec and zscaler-gre |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "jse": {
    "name": "name6",
    "num_users": 186,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "zscaler": {
    "aup_acceptance_required": false,
    "aup_expire": 210,
    "aup_ssl_proxy": false,
    "download_mbps": 94,
    "enable_aup": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

