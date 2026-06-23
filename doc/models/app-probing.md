
# App Probing

Application reachability probing settings for gateway management

## Structure

`AppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | Application keys from [List Applications](../../doc/controllers/constants-definitions.md#list-applications) |
| `CustomApps` | [`[]models.AppProbingCustomApp`](../../doc/models/app-probing-custom-app.md) | Optional | User-defined application probe definitions |
| `Enabled` | `*bool` | Optional | Whether gateway application probing is enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    appProbing := models.AppProbing{
        Apps:                 []string{
            "facebook",
        },
        CustomApps:           []models.AppProbingCustomApp{
            models.AppProbingCustomApp{
                Address:              models.ToPointer("address4"),
                AppType:              models.ToPointer("app_type2"),
                Hostnames:            []string{
                    "hostnames4",
                    "hostnames5",
                },
                Key:                  models.ToPointer("key8"),
                Name:                 models.ToPointer("name8"),
            },
            models.AppProbingCustomApp{
                Address:              models.ToPointer("address4"),
                AppType:              models.ToPointer("app_type2"),
                Hostnames:            []string{
                    "hostnames4",
                    "hostnames5",
                },
                Key:                  models.ToPointer("key8"),
                Name:                 models.ToPointer("name8"),
            },
        },
        Enabled:              models.ToPointer(false),
    }

}
```

