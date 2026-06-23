
# Map

Map or floorplan metadata for a site or organization

*This model accepts additional fields of type interface{}.*

## Structure

`Map`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Flags` | `map[string]int` | Optional, Read-only | Name/val pair objects for location engine to use |
| `ForSite` | `*bool` | Optional, Read-only | Whether this map belongs to a site scope |
| `Geofences` | [`[]models.MapGeofence`](../../doc/models/map-geofence.md) | Optional | List of geofences for the map |
| `GroupIdx` | `*int` | Optional | Optional floor or group ordering index for this map |
| `GroupName` | `*string` | Optional | Optional floor or group display name for this map |
| `Height` | `*int` | Optional | When `type`==`image`, height of the map image in pixels |
| `HeightM` | `*float64` | Optional | Physical height of the map in meters |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LatlngBr` | [`*models.LatlngBr`](../../doc/models/latlng-br.md) | Optional | When `type`==`google`, latitude and longitude of the bottom-right corner |
| `LatlngTl` | [`*models.LatlngTl`](../../doc/models/latlng-tl.md) | Optional | When `type`==`google`, latitude and longitude of the top-left corner |
| `Locked` | `*bool` | Optional | Whether this map is considered locked down<br><br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | The name of the map |
| `OccupancyLimit` | `*int` | Optional | Maximum occupancy configured for this map |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Orientation` | `*int` | Optional | Map orientation in degrees; 0 means up is north, 90 means up is west<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 359` |
| `OriginX` | `*int` | Optional | User-annotated X origin, pixels |
| `OriginY` | `*int` | Optional | User-annotated Y origin, pixels |
| `Ppm` | `*float64` | Optional | When `type`==`image`, pixels per meter for the map image |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SitesurveyPath` | [`[]models.MapSitesurveyPathItems`](../../doc/models/map-sitesurvey-path-items.md) | Optional | Site survey path definitions for a map<br><br>**Constraints**: *Minimum Items*: `0` |
| `ThumbnailUrl` | `*string` | Optional, Read-only | When `type`==`image`, URL for the thumbnail image or preview |
| `Type` | [`*models.MapTypeEnum`](../../doc/models/map-type-enum.md) | Optional | Map type, such as `image` or `google`. enum: `google`, `image`<br><br>**Default**: `"image"` |
| `Url` | `*string` | Optional, Read-only | When `type`==`image`, URL for the map image |
| `View` | [`models.Optional[models.MapViewEnum]`](../../doc/models/map-view-enum.md) | Optional | if `type`==`google`. enum: `hybrid`, `roadmap`, `satellite`, `terrain` |
| `WallPath` | [`*models.MapWallPath`](../../doc/models/map-wall-path.md) | Optional | JSON blob for wall definition (same format as wayfinding_path) |
| `Wayfinding` | [`*models.MapWayfinding`](../../doc/models/map-wayfinding.md) | Optional | Properties related to wayfinding |
| `WayfindingPath` | [`*models.MapWayfindingPath`](../../doc/models/map-wayfinding-path.md) | Optional | JSON blob for wayfinding (using Dijkstra’s algorithm) |
| `Width` | `*int` | Optional | When `type`==`image`, width of the map image in pixels |
| `WidthM` | `*float64` | Optional | Physical width of the map in meters |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mMap := models.Map{
        CreatedTime:          models.ToPointer(float64(212.24)),
        Flags:                map[string]int{
            "assetHoldTime": 5,
            "storeTime": 10,
        },
        ForSite:              models.ToPointer(false),
        Geofences:            []models.MapGeofence{
            models.MapGeofence{
                Name:                 models.ToPointer("name6"),
                Vertices:             []models.MapGeofenceVertice{
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                },
            },
            models.MapGeofence{
                Name:                 models.ToPointer("name6"),
                Vertices:             []models.MapGeofenceVertice{
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                },
            },
            models.MapGeofence{
                Name:                 models.ToPointer("name6"),
                Vertices:             []models.MapGeofenceVertice{
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                    models.MapGeofenceVertice{
                        X:                    models.ToPointer(float64(86.66)),
                        Y:                    models.ToPointer(float64(252.2)),
                    },
                },
            },
        },
        GroupIdx:             models.ToPointer(1),
        GroupName:            models.ToPointer("East Wing"),
        Height:               models.ToPointer(1500),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Locked:               models.ToPointer(false),
        Name:                 models.ToPointer("Mist Office"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Orientation:          models.ToPointer(30),
        OriginX:              models.ToPointer(35),
        OriginY:              models.ToPointer(60),
        Ppm:                  models.ToPointer(float64(40.94)),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        ThumbnailUrl:         models.ToPointer("https://url/to/image.png"),
        Type:                 models.ToPointer(models.MapTypeEnum_IMAGE),
        Url:                  models.ToPointer("https://url/to/image.png"),
        Width:                models.ToPointer(1250),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

