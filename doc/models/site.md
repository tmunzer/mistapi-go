
# Site

Site configuration and metadata within an organization

*This model accepts additional fields of type interface{}.*

## Structure

`Site`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `models.Optional[string]` | Optional | full address of the site |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | Alarm Template ID, this takes precedence over the Org-level alarmtemplate_id |
| `AptemplateId` | `models.Optional[uuid.UUID]` | Optional | AP Template ID, used by APs |
| `CountryCode` | `*string` | Optional | Country code for the site (for AP config generation), in two-character |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `GatewaytemplateId` | `models.Optional[uuid.UUID]` | Optional | Gateway Template ID, used by gateways |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | Geographic latitude and longitude coordinate pair |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the site |
| `NetworktemplateId` | `models.Optional[uuid.UUID]` | Optional | Network Template ID, this takes precedence over Site Settings |
| `Notes` | `models.Optional[string]` | Optional | Optional, any notes about the site |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RftemplateId` | `models.Optional[uuid.UUID]` | Optional | RF Template ID, this takes precedence over Site Settings |
| `RoutertemplateId` | `models.Optional[uuid.UUID]` | Optional | Router Template ID, used by gateways |
| `SecpolicyId` | `models.Optional[uuid.UUID]` | Optional | Security policy identifier applied to this site |
| `SitegroupIds` | `[]uuid.UUID` | Optional | Site group identifiers for groups this site belongs to |
| `SitetemplateId` | `models.Optional[uuid.UUID]` | Optional | Site template identifier applied to this site |
| `Timezone` | `*string` | Optional | IANA time zone name for the site<br><br>**Default**: `"UTC"` |
| `Tzoffset` | `*int` | Optional | Time zone offset value derived from the site's timezone<br><br>**Default**: `0` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    site := models.Site{
        Address:              models.NewOptional(models.ToPointer("1601 S. Deanza Blvd., Cupertino, CA, 95014")),
        AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("684dfc5c-fe77-2290-eb1d-ef3d677fe168"))),
        AptemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("16bdf952-ade2-4491-80b0-85ce506c760b"))),
        CountryCode:          models.ToPointer("US"),
        CreatedTime:          models.ToPointer(float64(113.34)),
        GatewaytemplateId:    models.NewOptional(models.ToPointer(uuid.MustParse("6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f"))),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 "Mist Office",
        NetworktemplateId:    models.NewOptional(models.ToPointer(uuid.MustParse("12ae9bd2-e0ab-107b-72e8-a7a005565ec2"))),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RftemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("bb8a9017-1e36-5d6c-6f2b-551abe8a76a2"))),
        RoutertemplateId:     models.NewOptional(models.ToPointer(uuid.MustParse("6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f"))),
        SecpolicyId:          models.NewOptional(models.ToPointer(uuid.MustParse("3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef"))),
        Timezone:             models.ToPointer("America/Los_Angeles"),
        Tzoffset:             models.ToPointer(0),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

