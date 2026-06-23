
# Site Setting Paloalto Network Gateway

Palo Alto Networks gateway API connection settings

## Structure

`SiteSettingPaloaltoNetworkGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiKey` | `*string` | Optional | Authentication key used to access the Palo Alto Networks gateway API |
| `ApiUrl` | `*string` | Optional | Base URL for the Palo Alto Networks gateway API |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingPaloaltoNetworkGateway := models.SiteSettingPaloaltoNetworkGateway{
        ApiKey:               models.ToPointer("5abf7c8a-1a1c-4398-ba2d-b0c297094d1a"),
        ApiUrl:               models.ToPointer("https://23.43.12.78:8443"),
    }

}
```

