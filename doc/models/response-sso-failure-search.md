
# Response Sso Failure Search

Response containing recent SSO authentication failure records

## Structure

`ResponseSsoFailureSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseSsoFailureSearchItem`](../../doc/models/response-sso-failure-search-item.md) | Required | SSO authentication failure records returned by the request<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsoFailureSearch := models.ResponseSsoFailureSearch{
        Results:              []models.ResponseSsoFailureSearchItem{
            models.ResponseSsoFailureSearchItem{
                Detail:               "detail2",
                SamlAssertionXml:     "saml_assertion_xml0",
                Timestamp:            float64(2.64),
            },
        },
    }

}
```

