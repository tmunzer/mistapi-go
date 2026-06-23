
# Mapstack

Map Stack filter or creation payload

## Structure

`Mapstack`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | The name of the map stack |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapstack := models.Mapstack{
        Name:                 models.ToPointer("Board Room"),
    }

}
```

