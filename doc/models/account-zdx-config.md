
# Account Zdx Config

OAuth linked ZDX apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZdxConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | ZDX cloud name. Refer https://help.zscaler.com/zdx/getting-started-zdx-api for ZDX cloud name<br><br>**Default**: `"zdxcloud.net"` |
| `KeyId` | `string` | Required | Customer account API key ID |
| `KeySecret` | `string` | Required | Customer account API key Secret |
| `ZdxOrgId` | `string` | Required | Organization identifier assigned to the ZDX account |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountZdxConfig := models.AccountZdxConfig{
        CloudName:            models.ToPointer("zdxcloud.net"),
        KeyId:                "K35vrZcK3JvrZc",
        KeySecret:            "K35vrZcK3JvrZcjjswpp2eii1oo100",
        ZdxOrgId:             "123456",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

