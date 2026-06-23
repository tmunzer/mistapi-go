
# Response Config History Search

Paginated device config history search response

## Structure

`ResponseConfigHistorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Search window end timestamp for config history, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of config history entries requested |
| `Next` | `*string` | Optional | URL for the next page of config history results |
| `Results` | [`[]models.ResponseConfigHistorySearchItem`](../../doc/models/response-config-history-search-item.md) | Required | Device config history entries returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Search window start timestamp for config history, in epoch seconds |
| `Total` | `int` | Required | Number of config history entries matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseConfigHistorySearch := models.ResponseConfigHistorySearch{
        End:                  44,
        Limit:                126,
        Next:                 models.ToPointer("next0"),
        Results:              []models.ResponseConfigHistorySearchItem{
            models.ResponseConfigHistorySearchItem{
                Channel24:            140,
                Channel5:             208,
                RadioMacs:            []string{
                    "radio_macs3",
                    "radio_macs4",
                    "radio_macs5",
                },
                Radios:               []models.ResponseConfigHistorySearchItemRadio{
                    models.ResponseConfigHistorySearchItemRadio{
                        Band:                 "band2",
                        Channel:              156,
                    },
                    models.ResponseConfigHistorySearchItemRadio{
                        Band:                 "band2",
                        Channel:              156,
                    },
                    models.ResponseConfigHistorySearchItemRadio{
                        Band:                 "band2",
                        Channel:              156,
                    },
                },
                SecpolicyViolated:    false,
                Ssids:                []string{
                    "ssids3",
                    "ssids4",
                },
                Ssids24:              []string{
                    "ssids_248",
                },
                Ssids5:               []string{
                    "ssids_56",
                    "ssids_57",
                },
                Timestamp:            float64(2.64),
                Version:              "version2",
            },
        },
        Start:                2,
        Total:                220,
    }

}
```

