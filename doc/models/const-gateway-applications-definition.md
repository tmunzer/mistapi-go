
# Const Gateway Applications Definition

Gateway application definition returned by the constants API

## Structure

`ConstGatewayApplicationsDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppId` | `*bool` | Optional | Whether an AppID is defined for this gateway application |
| `Key` | `*string` | Optional | Machine-readable key that identifies the gateway application |
| `Name` | `*string` | Optional | Display name of the gateway application |
| `SsrAppId` | `*bool` | Optional | Whether an SSR AppID is defined for this gateway application |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constGatewayApplicationsDefinition := models.ConstGatewayApplicationsDefinition{
        AppId:                models.ToPointer(true),
        Key:                  models.ToPointer("4shared"),
        Name:                 models.ToPointer("4shared"),
        SsrAppId:             models.ToPointer(true),
    }

}
```

