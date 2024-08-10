
# Template

Template

## Structure

`Template`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.TemplateApplies`](../../doc/models/template-applies.md) | Optional | where this template should be applied to, can be org_id, site_ids, sitegroup_ids |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileIds` | `[]uuid.UUID` | Optional | list of Device Profile ids |
| `Exceptions` | [`*models.TemplateExceptions`](../../doc/models/template-exceptions.md) | Optional | where this template should not be applied to (takes precedence) |
| `FilterByDeviceprofile` | `*bool` | Optional | whether to further filter by Device Profile |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name6",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "applies": {
    "org_id": "00001bdc-0000-0000-0000-000000000000",
    "site_ids": [
      "00001706-0000-0000-0000-000000000000",
      "00001707-0000-0000-0000-000000000000",
      "00001708-0000-0000-0000-000000000000"
    ],
    "sitegroup_ids": [
      "00000634-0000-0000-0000-000000000000"
    ]
  },
  "created_time": 247.46,
  "deviceprofile_ids": [
    "0000104d-0000-0000-0000-000000000000",
    "0000104e-0000-0000-0000-000000000000",
    "0000104f-0000-0000-0000-000000000000"
  ],
  "exceptions": {
    "site_ids": [
      "00001604-0000-0000-0000-000000000000",
      "00001603-0000-0000-0000-000000000000",
      "00001602-0000-0000-0000-000000000000"
    ],
    "sitegroup_ids": [
      "00000c2e-0000-0000-0000-000000000000"
    ]
  },
  "filter_by_deviceprofile": false
}
```

