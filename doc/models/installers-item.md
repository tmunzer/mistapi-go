
# Installers Item

Installer access entry with scope and role details

## Structure

`InstallersItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | Display name of the installer access entry |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    installersItem := models.InstallersItem{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("Entry #1"),
    }

}
```

