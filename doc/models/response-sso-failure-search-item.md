
# Response Sso Failure Search Item

SSO authentication failure record

## Structure

`ResponseSsoFailureSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | Failure details reported for the SSO authentication attempt |
| `SamlAssertionXml` | `string` | Required | SAML assertion XML captured for the failed SSO authentication attempt |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsoFailureSearchItem := models.ResponseSsoFailureSearchItem{
        Detail:               "detail0",
        SamlAssertionXml:     "saml_assertion_xml2",
        Timestamp:            float64(88.88),
    }

}
```

