
# Sdktemplate

SDK Template

## Structure

`Sdktemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BgImage` | `*string` | Optional | - |
| `BtnFlrBgcolor` | `*string` | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `Default` | `*bool` | Optional | whether this is the default template when there are multiple templates |
| `ForSite` | `*bool` | Optional | - |
| `HeaderTxt` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | name for identification purpose |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SearchTxtcolor` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `WelcomeMsg` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "bg_image": "bg_image0",
  "btn_flr_bgcolor": "btn_flr_bgcolor4",
  "created_time": 225.64,
  "default": false,
  "for_site": false
}
```

