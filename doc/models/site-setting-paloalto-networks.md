
# Site Setting Paloalto Networks

## Structure

`SiteSettingPaloaltoNetworks`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateways` | [`[]models.SiteSettingPaloaltoNetworkGateway`](../../doc/models/site-setting-paloalto-network-gateway.md) | Optional | - |
| `SendMistNacUserInfo` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "send_mist_nac_user_info": false,
  "gateways": [
    {
      "api_key": "api_key8",
      "api_url": "api_url0"
    }
  ]
}
```

