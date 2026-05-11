
# Account Skyatp Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSkyatpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | [`*models.AccountSkyatpCloudNameEnum`](../../doc/models/account-skyatp-cloud-name-enum.md) | Optional | Sky ATP cloud name. enum: `www.amerskyatp.com`, `www.apacskyatp.com`, `www.euroskyatp.com`, `www.canadaskyatp.com` |
| `Password` | `string` | Required | - |
| `Realm` | `string` | Required | - |
| `Username` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "www.amerskyatp.com",
  "password": "foryoureyesonly",
  "realm": "mist-team",
  "username": "john@abc.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

