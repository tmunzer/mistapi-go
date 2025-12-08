
# Site Setting Juniper Srx

## Structure

`SiteSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.JuniperSrxAutoUpgrade`](../../doc/models/juniper-srx-auto-upgrade.md) | Optional | auto_upgrade device first time it is onboarded |
| `Gateways` | [`[]models.SiteSettingJuniperSrxGateway`](../../doc/models/site-setting-juniper-srx-gateway.md) | Optional | - |
| `SendMistNacUserInfo` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "auto_upgrade": {
    "custom_versions": {
      "key0": "custom_versions3",
      "key1": "custom_versions2"
    },
    "enabled": false,
    "snapshot": false
  },
  "gateways": [
    {
      "api_key": "api_key8",
      "api_password": "api_password8",
      "api_url": "api_url0"
    },
    {
      "api_key": "api_key8",
      "api_password": "api_password8",
      "api_url": "api_url0"
    }
  ],
  "send_mist_nac_user_info": false
}
```

