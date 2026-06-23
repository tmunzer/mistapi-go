
# Installer Site

Site definition available to installer workflows

*This model accepts additional fields of type interface{}.*

## Structure

`InstallerSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `string` | Required | Street or postal address for the installer site |
| `CountryCode` | `string` | Required | ISO country code for the installer site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | Geographic latitude and longitude coordinate pair |
| `Name` | `string` | Required | Display name of the installer site |
| `RftemplateName` | `*string` | Optional | RF template name applied to the installer site |
| `SitegroupNames` | `[]string` | Optional | Site group names associated with an installer site |
| `Timezone` | `*string` | Optional | Time zone configured for the installer site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    installerSite := models.InstallerSite{
        Address:              "1601 S. Deanza Blvd., Cupertino, CA, 95014",
        CountryCode:          "US",
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Latlng:               models.LatLng{
            Lat:                  float64(37.295833),
            Lng:                  float64(-122.032946),
        },
        Name:                 "Mist Office",
        RftemplateName:       models.ToPointer("rftemplate1"),
        SitegroupNames:       []string{
            "sg1",
            "sg2",
        },
        Timezone:             models.ToPointer("America/Los_Angeles"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

