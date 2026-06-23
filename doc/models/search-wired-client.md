
# Search Wired Client

Paginated response for wired client searches

## Structure

`SearchWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Upper bound timestamp of the wired client search window, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of wired client records returned by this page |
| `Next` | `*string` | Optional | URL for the next page of wired client records, when more results are available |
| `Results` | [`[]models.WiredClientResponse`](../../doc/models/wired-client-response.md) | Required | Wired client records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Lower bound timestamp of the wired client search window, in epoch seconds |
| `Total` | `int` | Required | Count of wired client records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    searchWiredClient := models.SearchWiredClient{
        End:                  float64(233.04),
        Limit:                162,
        Next:                 models.ToPointer("next6"),
        Results:              []models.WiredClientResponse{
            models.WiredClientResponse{
                AuthMethod:                models.ToPointer("mac_auth"),
                AuthState:                 models.ToPointer("authenticated"),
                DeviceMac:                 []string{
                    "device_mac7",
                },
                DeviceMacPort:             []models.WiredClientResponseDeviceMacPortItem{
                    models.WiredClientResponseDeviceMacPortItem{
                        DeviceMac:            models.ToPointer("device_mac8"),
                        Ip:                   models.ToPointer("ip8"),
                        PortId:               models.ToPointer("port_id4"),
                        PortParent:           models.ToPointer("port_parent6"),
                        Start:                models.ToPointer("start8"),
                    },
                    models.WiredClientResponseDeviceMacPortItem{
                        DeviceMac:            models.ToPointer("device_mac8"),
                        Ip:                   models.ToPointer("ip8"),
                        PortId:               models.ToPointer("port_id4"),
                        PortParent:           models.ToPointer("port_parent6"),
                        Start:                models.ToPointer("start8"),
                    },
                },
                DhcpClientIdentifier:      models.ToPointer("MAC address 00155df6d500"),
                DhcpFqdn:                  models.ToPointer("ITS-VMMT0-D1N02.mgthub.local"),
                DhcpHostname:              models.ToPointer("ITS-VMMT0-D1N02"),
                DhcpRequestParams:         models.ToPointer("1 3 6 15 31 33 43 44 46 47 119 121 249 252"),
                DhcpVendorClassIdentifier: models.ToPointer("MSFT 5.0"),
                OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Start:                float64(189.1),
        Total:                0,
    }

}
```

