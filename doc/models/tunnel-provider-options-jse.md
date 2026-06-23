
# Tunnel Provider Options Jse

For jse-ipsec, this allows provisioning of adequate resource on JSE. Make sure adequate licenses are added

## Structure

`TunnelProviderOptionsJse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumUsers` | `*int` | Optional | User capacity to provision on Juniper Secure Edge |
| `OrgName` | `*string` | Optional | JSE Organization name. The list of available organizations can be retrieved with the [Get Org JSE Info]($e/Orgs%20JSE/getOrgJseInfo) API Call |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelProviderOptionsJse := models.TunnelProviderOptionsJse{
        NumUsers:             models.ToPointer(5),
        OrgName:              models.ToPointer("JSE_ORG1"),
    }

}
```

