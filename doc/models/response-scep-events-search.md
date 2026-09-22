
# Response Scep Events Search

Paginated Mist SCEP PKI operation event search response

## Structure

`ResponseScepEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | End of the SCEP event search window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of SCEP events returned per page |
| `Page` | `*int` | Optional | Current page of SCEP event search results |
| `Results` | [`[]models.ScepEvent`](../../doc/models/scep-event.md) | Optional | SCEP PKI operation events returned by a search |
| `Start` | `*int` | Optional | Start of the SCEP event search window, in epoch seconds |
| `Total` | `*int` | Optional | Number of SCEP events matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseScepEventsSearch := models.ResponseScepEventsSearch{
        End:                  models.ToPointer(1748314800),
        Limit:                models.ToPointer(100),
        Page:                 models.ToPointer(1),
        Results:              []models.ScepEvent{
            models.ScepEvent{
                CertProvider:         models.ToPointer("cert_provider6"),
                CommonName:           models.ToPointer("common_name4"),
                DeviceId:             models.ToPointer("device_id2"),
                Text:                 models.ToPointer("text4"),
            },
            models.ScepEvent{
                CertProvider:         models.ToPointer("cert_provider6"),
                CommonName:           models.ToPointer("common_name4"),
                DeviceId:             models.ToPointer("device_id2"),
                Text:                 models.ToPointer("text4"),
            },
        },
        Start:                models.ToPointer(1748228400),
        Total:                models.ToPointer(3),
    }

}
```

