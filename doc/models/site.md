
# Site

Site

*This model accepts additional fields of type interface{}.*

## Structure

`Site`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | full address of the site |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | Alarm Template ID, this takes precedence over the Org-level alarmtemplate_id |
| `AptemplateId` | `models.Optional[uuid.UUID]` | Optional | AP Template ID, used by APs |
| `CountryCode` | `*string` | Optional | Country code for the site (for AP config generation), in two-character |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `GatewaytemplateId` | `models.Optional[uuid.UUID]` | Optional | Gateway Template ID, used by gateways |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `NetworktemplateId` | `models.Optional[uuid.UUID]` | Optional | Network Template ID, this takes precedence over Site Settings |
| `Notes` | `*string` | Optional | Optional, any notes about the site |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RftemplateId` | `models.Optional[uuid.UUID]` | Optional | RF Template ID, this takes precedence over Site Settings |
| `SecpolicyId` | `models.Optional[uuid.UUID]` | Optional | SecPolicy ID |
| `SitegroupIds` | `[]uuid.UUID` | Optional | Sitegroups this site belongs to |
| `SitetemplateId` | `models.Optional[uuid.UUID]` | Optional | Site Template ID |
| `Timezone` | `*string` | Optional | Timezone the site is at<br>**Default**: `"UTC"` |
| `Tzoffset` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
  "country_code": "US",
  "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Mist Office",
  "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
  "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
  "timezone": "America/Los_Angeles",
  "created_time": 113.34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

