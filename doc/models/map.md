
# Map

Map

## Structure

`Map`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Flags` | `map[string]int` | Optional | name/val pair objects for location engine to use |
| `ForSite` | `*bool` | Optional | - |
| `Height` | `*int` | Optional | when type=image, height of the image map |
| `HeightM` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `LatlngBr` | [`*models.LatlngBr`](../../doc/models/latlng-br.md) | Optional | when type=google, latitude / longitude of the bottom-right corner |
| `LatlngTl` | [`*models.LatlngTl`](../../doc/models/latlng-tl.md) | Optional | when type=google, latitude / longitude of the top-left corner |
| `Locked` | `*bool` | Optional | whether this map is considered locked down<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | The name of the map |
| `OccupancyLimit` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Orientation` | `*int` | Optional | orientation of the map, 0 means up is north, 90 means up is west<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 359` |
| `OriginX` | `*int` | Optional | the user-annotated x origin, pixels |
| `OriginY` | `*int` | Optional | the user-annotated y origin, pixels |
| `Ppm` | `*float64` | Optional | when type=image, pixels per meter |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SitesurveyPath` | [`[]models.MapSitesurveyPathItems`](../../doc/models/map-sitesurvey-path-items.md) | Optional | sitesurvey_path<br>**Constraints**: *Minimum Items*: `0` |
| `ThumbnailUrl` | `*string` | Optional | when type=image, the url for the thumbnail image / preview |
| `Type` | [`*models.MapTypeEnum`](../../doc/models/map-type-enum.md) | Optional | enum: `google`, `image`<br>**Default**: `"image"` |
| `Url` | `*string` | Optional | when type=image, the url |
| `UseAutoOrientation` | `*bool` | Optional | whether this map uses autooreintation values or ignores them<br>**Default**: `false` |
| `UseAutoPlacement` | `*bool` | Optional | whether this map uses autoplacement values or ignores them<br>**Default**: `false` |
| `View` | [`models.Optional[models.MapViewEnum]`](../../doc/models/map-view-enum.md) | Optional | if `type`==`google`. enum: `hybrid`, `roadmap`, `satellite`, `terrain` |
| `WallPath` | [`*models.MapWallPath`](../../doc/models/map-wall-path.md) | Optional | a JSON blob for wall definition (same format as wayfinding_path) |
| `Wayfinding` | [`*models.MapWayfinding`](../../doc/models/map-wayfinding.md) | Optional | properties related to wayfinding |
| `WayfindingPath` | [`*models.MapWayfindingPath`](../../doc/models/map-wayfinding-path.md) | Optional | a JSON blob for wayfinding (using Dijkstra’s algorithm) |
| `Width` | `*int` | Optional | when type=image, width of the image map |
| `WidthM` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "flags": {
    "assetHoldTime": 5,
    "storeTime": 10
  },
  "height": 1500,
  "locked": false,
  "name": "Mist Office",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "orientation": 30,
  "origin_x": 35,
  "origin_y": 60,
  "ppm": 40.94,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "thumbnail_url": "https://url/to/image.png",
  "type": "image",
  "url": "https://url/to/image.png",
  "use_auto_orientation": false,
  "use_auto_placement": false,
  "width": 1250,
  "created_time": 45.28,
  "for_site": false,
  "height_m": 113.3
}
```

