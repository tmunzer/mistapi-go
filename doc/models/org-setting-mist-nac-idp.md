
# Org Setting Mist Nac Idp

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingMistNacIdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExcludeRealms` | `[]string` | Optional | When the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `UserRealms` | `[]string` | Optional | Which realm should trigger this IDP. User Realm is extracted from:<br><br>* Username-AVP (`mist.com` from john@mist.com)<br>* Cert CN |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

