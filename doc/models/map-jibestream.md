
# Map Jibestream

Jibestream map import configuration

## Structure

`MapJibestream`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `uuid.UUID` | Required | Client identifier for the Jibestream API |
| `ClientSecret` | `string` | Required | Client secret for the Jibestream API |
| `CustomerId` | `int` | Required | Jibestream customer record id |
| `EndpointUrl` | `string` | Required | Map contents endpoint host |
| `MapId` | `uuid.UUID` | Required | Jibestream map identifier to import |
| `Mmpp` | `int` | Required | Millimeters per pixel for the Jibestream map |
| `Ppm` | `float64` | Required | Pixel per meter, same as the map JSON value. |
| `VendorName` | `string` | Required, Constant | The vendor ‘jibestream’. enum: `jibestream`<br><br>**Value**: `"jibestream"` |
| `VenueId` | `int` | Required | Venue or organization id |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mapJibestream := models.MapJibestream{
        ClientId:             uuid.MustParse("199d6770-0f6f-407a-9bd5-fc33c7840194"),
        ClientSecret:         "/9Nog3yDzcYj0bY91XJZQLCt+m9DXaIVhx+Ghk3ddd",
        CustomerId:           123,
        EndpointUrl:          "https://api.jibestream.com",
        MapId:                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
        Mmpp:                 223,
        Ppm:                  float64(4),
        VendorName:           "jibestream",
        VenueId:              123,
    }

}
```

