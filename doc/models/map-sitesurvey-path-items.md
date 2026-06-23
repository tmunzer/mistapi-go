
# Map Sitesurvey Path Items

Site survey path definition on a map

## Structure

`MapSitesurveyPathItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | Site survey path coordinate space |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | Display name of the site survey path |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | Map nodes included in a site survey path<br><br>**Constraints**: *Minimum Items*: `0` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mapSitesurveyPathItems := models.MapSitesurveyPathItems{
        Coordinate:           models.ToPointer("actual"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("Default"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "key0": "edges6",
                },
                Name:                 "name6",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(224.72),
                    Y:                    float64(100),
                }),
            },
            models.MapNode{
                Edges:                map[string]string{
                    "key0": "edges6",
                },
                Name:                 "name6",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(224.72),
                    Y:                    float64(100),
                }),
            },
            models.MapNode{
                Edges:                map[string]string{
                    "key0": "edges6",
                },
                Name:                 "name6",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(224.72),
                    Y:                    float64(100),
                }),
            },
        },
    }

}
```

