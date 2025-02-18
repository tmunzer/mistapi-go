
# Sdktemplate

SDK Template

*This model accepts additional fields of type interface{}.*

## Structure

`Sdktemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BgImage` | `*string` | Optional | - |
| `BtnFlrBgcolor` | `*string` | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Default` | `*bool` | Optional | Whether this is the default template when there are multiple templates |
| `ForSite` | `*bool` | Optional | - |
| `HeaderTxt` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Name for identification purpose |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SearchTxtcolor` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `WelcomeMsg` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "bg_image": "bg_image0",
  "btn_flr_bgcolor": "btn_flr_bgcolor4",
  "created_time": 41.34,
  "default": false,
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

