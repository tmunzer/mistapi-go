
# Template

Template

*This model accepts additional fields of type interface{}.*

## Structure

`Template`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.TemplateApplies`](../../doc/models/template-applies.md) | Optional | Where this template should be applied to, can be org_id, site_ids, sitegroup_ids |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileIds` | `[]uuid.UUID` | Optional | List of Device Profile ids |
| `Exceptions` | [`*models.TemplateExceptions`](../../doc/models/template-exceptions.md) | Optional | Where this template should not be applied to (takes precedence) |
| `FilterByDeviceprofile` | `*bool` | Optional | Whether to further filter by Device Profile |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
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
  "filter_by_deviceprofile": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

