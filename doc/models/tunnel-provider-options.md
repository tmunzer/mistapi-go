
# Tunnel Provider Options

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Jse` | [`*models.TunnelProviderOptionsJse`](../../doc/models/tunnel-provider-options-jse.md) | Optional | For jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added |
| `Zscaler` | [`*models.TunnelProviderOptionsZscaler`](../../doc/models/tunnel-provider-options-zscaler.md) | Optional | For zscaler-ipsec and zscaler-gre |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "jse": {
    "num_users": 186,
    "org_name": "org_name6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "zscaler": {
    "aup_block_internet_until_accepted": false,
    "aup_enabled": false,
    "aup_force_ssl_inspection": false,
    "aup_timeout_in_days": 104,
    "auth_required": false,
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

