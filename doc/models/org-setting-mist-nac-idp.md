
# Org Setting Mist Nac Idp

## Structure

`OrgSettingMistNacIdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExcludeRealms` | `[]string` | Optional | when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org |
| `Id` | `*uuid.UUID` | Optional | - |
| `UserRealms` | `[]string` | Optional | which realm should trigger this IDP. User Realm is extracted from:<br><br>* Username-AVP (`mist.com` from john@mist.com)<br>* Cert CN |

## Example (as JSON)

```json
{
  "id": "4c441a74-d0de-32c4-78a7-a05e00d080ae",
  "exclude_realms": [
    "exclude_realms9"
  ],
  "user_realms": [
    "user_realms2",
    "user_realms3"
  ]
}
```

