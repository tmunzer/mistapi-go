
# Response Guest Search

Paginated response for guest authorization search results

## Structure

`ResponseGuestSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the guest authorization search window |
| `Limit` | `int` | Required | Maximum number of guest authorization records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of guest authorization records |
| `Results` | [`[]models.Guest`](../../doc/models/guest.md) | Required | Guest authorization records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the guest authorization search window |
| `Total` | `int` | Required | Number of guest authorization records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseGuestSearch := models.ResponseGuestSearch{
        End:                  94,
        Limit:                76,
        Next:                 models.ToPointer("next0"),
        Results:              []models.Guest{
            models.Guest{
                AccessCodeEmail:        models.ToPointer("access_code_email8"),
                ApMac:                  models.ToPointer("ap_mac8"),
                AuthMethod:             models.ToPointer("auth_method0"),
                Authorized:             models.ToPointer(true),
                AuthorizedExpiringTime: models.ToPointer(float64(1480704955)),
                AuthorizedTime:         models.ToPointer(float64(1480704355)),
                Company:                models.ToPointer("abc"),
                Email:                  models.ToPointer("john@abc.com"),
                Minutes:                models.ToPointer(1440),
                Name:                   models.ToPointer("John Smith"),
                Ssid:                   models.ToPointer("Guest-SSID"),
                WlanId:                 models.ToPointer(uuid.MustParse("6748cfa6-4e12-11e6-9188-0242ac110007")),
                AdditionalProperties:   map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        Start:                52,
        Total:                170,
    }

}
```

