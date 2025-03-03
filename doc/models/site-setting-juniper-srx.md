
# Site Setting Juniper Srx

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateways` | [`[]models.SiteSettingJuniperSrxGateway`](../../doc/models/site-setting-juniper-srx-gateway.md) | Optional | - |
| `SendMistNacUserInfo` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "gateways": [
    {
      "api_key": "api_key8",
      "api_password": "api_password8",
      "api_url": "api_url0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "api_key": "api_key8",
      "api_password": "api_password8",
      "api_url": "api_url0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "send_mist_nac_user_info": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

