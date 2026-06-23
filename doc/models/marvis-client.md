
# Marvis Client

Marvis Client configuration profile

*This model accepts additional fields of type interface{}.*

## Structure

`MarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Whether this Marvis Client profile is disabled<br><br>**Default**: `false` |
| `EnrollmentUrl` | `*string` | Optional, Read-only | In MDM, add `--enrollment_url <enrollment_url>` to the install command |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Location` | [`*models.MarvisClientLocation`](../../doc/models/marvis-client-location.md) | Optional | Location collection settings for Marvis Client |
| `Name` | `*string` | Optional | Display name for the Marvis Client profile |
| `SyntheticTest` | [`*models.MarvisClientSyntheticTest`](../../doc/models/marvis-client-synthetic-test.md) | Optional | Synthetic test settings for Marvis Client |
| `Telemetry` | [`*models.MarvisClientTelemetry`](../../doc/models/marvis-client-telemetry.md) | Optional | Note: some stats are not collected when it's not connected to Mist infrastructure |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    marvisClient := models.MarvisClient{
        Disabled:             models.ToPointer(false),
        EnrollmentUrl:        models.ToPointer("marvisclient://api.mist.com/path/to/url"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Location:             models.ToPointer(models.MarvisClientLocation{
            Enabled:              models.ToPointer(false),
        }),
        Name:                 models.ToPointer("Handhelds"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

