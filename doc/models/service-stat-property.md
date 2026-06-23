
# Service Stat Property

Version information for gateway service packages

## Structure

`ServiceStatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AshVersion` | `*string` | Optional | Version of the ASH service package |
| `CiaVersion` | `*string` | Optional | Version of the CIA service package |
| `EmberVersion` | `*string` | Optional | Version of the Ember service package |
| `IpsecClientVersion` | `*string` | Optional | Version of the IPsec client package |
| `MistAgentVersion` | `*string` | Optional | Version of the Mist agent package |
| `PackageVersion` | `*string` | Optional | Version of the service package |
| `TestingToolsVersion` | `*string` | Optional | Version of the testing tools package |
| `WheeljackVersion` | `*string` | Optional | Version of the Wheeljack service package |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    serviceStatProperty := models.ServiceStatProperty{
        AshVersion:           models.ToPointer("ash_version8"),
        CiaVersion:           models.ToPointer("cia_version4"),
        EmberVersion:         models.ToPointer("ember_version0"),
        IpsecClientVersion:   models.ToPointer("ipsec_client_version0"),
        MistAgentVersion:     models.ToPointer("mist_agent_version2"),
    }

}
```

