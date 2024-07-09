
# Org Setting Mist Nac Idp

## Structure

`OrgSettingMistNacIdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExcludeRealms` | `[]string` | Optional | when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org |
| `Id` | `*string` | Optional | - |
| `UserRealms` | `[]string` | Optional | which realm should trigger this IDP.<br>we extract user realm from<br><br>1. Username-AVP (`mist.com` from john@mist.com)<br>2. Cert CN |

## Example (as JSON)

```json
{
  "id": "4c441a74-d0de-32c4-78a7-a05e00d080ae",
  "exclude_realms": [
    "exclude_realms5",
    "exclude_realms4"
  ],
  "user_realms": [
    "user_realms8"
  ]
}
```

