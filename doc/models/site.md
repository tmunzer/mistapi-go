
# Site

Site

## Structure

`Site`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | full address of the site |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | Alarm Template ID, this takes precedence over the Org-level alarmtemplate_id |
| `AptemplateId` | `models.Optional[uuid.UUID]` | Optional | AP Template ID, used by APs |
| `CountryCode` | `*string` | Optional | country code for the site (for AP config generation), in two-character |
| `CreatedTime` | `*float64` | Optional | - |
| `GatewaytemplateId` | `models.Optional[uuid.UUID]` | Optional | Gateway Template ID, used by gateways |
| `Id` | `*uuid.UUID` | Optional | - |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `NetworktemplateId` | `models.Optional[uuid.UUID]` | Optional | Network Template ID, this takes precedence over Site Settings |
| `Notes` | `*string` | Optional | optional, any notes about the site |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RftemplateId` | `models.Optional[uuid.UUID]` | Optional | RF Template ID, this takes precedence over Site Settings |
| `SecpolicyId` | `models.Optional[uuid.UUID]` | Optional | SecPolicy ID |
| `SitegroupIds` | `[]uuid.UUID` | Optional | sitegroups this site belongs to |
| `SitetemplateId` | `models.Optional[uuid.UUID]` | Optional | Site Template ID |
| `Timezone` | `*string` | Optional | Timezone the site is at |

## Example (as JSON)

```json
{
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
  "country_code": "US",
  "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
  "name": "Mist Office",
  "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
  "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
  "timezone": "America/Los_Angeles",
  "created_time": 154.0
}
```
