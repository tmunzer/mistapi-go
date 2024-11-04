
# Org Setting Mist Nac Idp

## Structure

`OrgSettingMistNacIdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExcludeRealms` | `[]string` | Optional | when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `UserRealms` | `[]string` | Optional | which realm should trigger this IDP. User Realm is extracted from:<br><br>* Username-AVP (`mist.com` from john@mist.com)<br>* Cert CN |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "exclude_realms": [
    "exclude_realms9"
  ],
  "user_realms": [
    "user_realms2",
    "user_realms3"
  ]
}
```

