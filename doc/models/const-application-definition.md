
# Const Application Definition

Application definition recognized by Juniper Devices

## Structure

`ConstApplicationDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppId` | `*bool` | Optional | Whether an AppID is defined for this application |
| `AppImageUrl` | `*string` | Optional | URL of the application icon image, when one is available |
| `AppProbe` | `*bool` | Optional | Whether an application probe is available for this application |
| `Category` | `*string` | Optional | Application category key associated with this application |
| `Group` | `*string` | Optional | Application group display name for this application |
| `Key` | `*string` | Optional | Machine-readable key that identifies the application |
| `Name` | `*string` | Optional | Display name of the application |
| `SignatureBased` | `*bool` | Optional | Whether this application is detected using signatures |
| `SsrAppId` | `*bool` | Optional | Whether an SSR AppID is defined for this application |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constApplicationDefinition := models.ConstApplicationDefinition{
        AppId:                models.ToPointer(false),
        AppImageUrl:          models.ToPointer("\"\""),
        AppProbe:             models.ToPointer(false),
        Category:             models.ToPointer("FileSharing"),
        Group:                models.ToPointer("File Sharing"),
        Key:                  models.ToPointer("dropbox"),
        Name:                 models.ToPointer("Dropbox"),
    }

}
```

