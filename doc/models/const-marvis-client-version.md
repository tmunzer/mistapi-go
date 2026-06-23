
# Const Marvis Client Version

Marvis Client version download entry

## Structure

`ConstMarvisClientVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Label` | `*string` | Optional | Release label for this Marvis Client version |
| `Notes` | `*string` | Optional | Release notes or extra text for this Marvis Client version |
| `Os` | `*string` | Optional | Operating system supported by this Marvis Client installer |
| `Url` | `*string` | Optional | Download URL for this Marvis Client installer |
| `Version` | `*string` | Optional | Marvis Client software version for this installer |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constMarvisClientVersion := models.ConstMarvisClientVersion{
        Label:                models.ToPointer("default"),
        Notes:                models.ToPointer("notes0"),
        Os:                   models.ToPointer("windows"),
        Url:                  models.ToPointer("https://mobile.mist.com/installers/marvisclient/..."),
        Version:              models.ToPointer("0.100.29"),
    }

}
```

