
# Account Cradlepoint Config

Cradlepoint account credentials and LLDP integration settings

*This model accepts additional fields of type interface{}.*

## Structure

`AccountCradlepointConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpApiId` | `*string` | Optional | Cradlepoint API ID used by Mist for the integration |
| `CpApiKey` | `*string` | Optional | Cradlepoint API key paired with the Cradlepoint API ID |
| `EcmApiId` | `*string` | Optional | Cradlepoint ECM API ID used by Mist for the integration |
| `EcmApiKey` | `*string` | Optional | Cradlepoint ECM API key paired with the ECM API ID |
| `EnableLldp` | `*bool` | Optional | Whether Mist uses Cradlepoint LLDP data to link routers to Mist sites and devices |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountCradlepointConfig := models.AccountCradlepointConfig{
        CpApiId:              models.ToPointer("84446d61-2206-4ea5-855a-0043f980be54"),
        CpApiKey:             models.ToPointer("79c329da9893e34099c7d8ad5cb9c941"),
        EcmApiId:             models.ToPointer("73446d61-2206-4ea5-855a-0043f980be62"),
        EcmApiKey:            models.ToPointer("68b329da9893e34099c7d8ad5cb9c9405"),
        EnableLldp:           models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

