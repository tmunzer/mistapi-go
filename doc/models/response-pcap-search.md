
# Response Pcap Search

Paginated response for packet capture search results

## Structure

`ResponsePcapSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the packet capture search window |
| `Limit` | `int` | Required | Maximum number of packet capture records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of packet capture search results |
| `Results` | [`[]models.ResponsePcapSearchItem`](../../doc/models/response-pcap-search-item.md) | Required | Packet capture records returned by search<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the packet capture search window |
| `Total` | `*int` | Optional | Number of packet capture records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePcapSearch := models.ResponsePcapSearch{
        End:                  254,
        Limit:                172,
        Next:                 models.ToPointer("next4"),
        Results:              []models.ResponsePcapSearchItem{
            models.ResponsePcapSearchItem{
                ApMacs:               []string{
                    "ap_macs9",
                    "ap_macs8",
                },
                Aps:                  []string{
                    "aps7",
                    "aps8",
                },
                Duration:             models.ToPointer(float64(600)),
                Format:               models.ToPointer("stream"),
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                LastSeen:             models.ToPointer(float64(1693482149.417)),
                MaxNumPackets:        models.ToPointer(float64(1024)),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                PcapAps:              map[string]models.ResponsePcapSearchItemPcapApsItem{
                    "5c5b35000010": models.ResponsePcapSearchItemPcapApsItem{
                        Band:                 models.ToPointer("6"),
                        Bandwidth:            models.ToPointer("20"),
                        Channel:              models.ToPointer(133),
                        TcpdumpExpression:    models.NewOptional[string](nil),
                    },
                },
                TerminationReason:    models.ToPointer("default"),
                Timestamp:            float64(2.64),
                Type:                 "type4",
                Url:                  "url0",
            },
        },
        Start:                212,
        Total:                models.ToPointer(246),
    }

}
```

