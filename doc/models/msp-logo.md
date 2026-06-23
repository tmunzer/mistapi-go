
# Msp Logo

Payload for uploading an advanced-tier MSP logo

*This model accepts additional fields of type interface{}.*

## Structure

`MspLogo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LogoUrl` | `*string` | Optional | Public URL for the advanced-tier MSP logo image |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mspLogo := models.MspLogo{
        LogoUrl:              models.ToPointer("logo_url2"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

