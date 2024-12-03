
# Site Setting Paloalto Networks

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingPaloaltoNetworks`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateways` | [`[]models.SiteSettingPaloaltoNetworkGateway`](../../doc/models/site-setting-paloalto-network-gateway.md) | Optional | - |
| `SendMistNacUserInfo` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "send_mist_nac_user_info": false,
  "gateways": [
    {
      "api_key": "api_key8",
      "api_url": "api_url0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

