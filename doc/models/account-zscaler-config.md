
# Account Zscaler Config

OAuth linked Zscaler apps account details

## Structure

`AccountZscalerConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `string` | Required | - |
| `PartnerKey` | `string` | Required | - |
| `Password` | `string` | Required | customer account password |
| `Username` | `string` | Required | customer account user name |

## Example (as JSON)

```json
{
  "cloud_name": "zscalerbeta.net",
  "partner_key": "K35vrZcK3JvrZc",
  "password": "password",
  "username": "john@nmo.com"
}
```

