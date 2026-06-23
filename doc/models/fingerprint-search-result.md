
# Fingerprint Search Result

Search response for client device fingerprint records

## Structure

`FingerprintSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Upper bound of the time range for the fingerprint search |
| `Limit` | `int` | Required | Maximum number of fingerprint records returned in the response |
| `Next` | `*string` | Optional | Pagination URL for the next page of fingerprint search results |
| `Results` | [`[]models.Fingerprint`](../../doc/models/fingerprint.md) | Required | Client device fingerprint records |
| `Start` | `int` | Required | Lower bound of the time range for the fingerprint search |
| `Total` | `int` | Required | Number of fingerprint records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    fingerprintSearchResult := models.FingerprintSearchResult{
        End:                  1711035686,
        Limit:                10,
        Next:                 models.ToPointer("next8"),
        Results:              []models.Fingerprint{
            models.Fingerprint{
                Family:               models.ToPointer("family8"),
                Mac:                  models.ToPointer("mac0"),
                Mfg:                  models.ToPointer("mfg6"),
                Model:                models.ToPointer("model4"),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Start:                1710949286,
        Total:                232,
    }

}
```

