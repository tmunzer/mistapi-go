
# Account Skyatp Info

## Structure

`AccountSkyatpInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | [`*models.AccountSkyatpCloudNameEnum`](../../doc/models/account-skyatp-cloud-name-enum.md) | Optional | Sky ATP cloud name. enum: `www.amerskyatp.com`, `www.apacskyatp.com`, `www.euroskyatp.com`, `www.canadaskyatp.com` |
| `Realm` | `*string` | Optional | - |
| `Username` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "www.amerskyatp.com",
  "realm": "mist-team",
  "username": "john@abc.com"
}
```

