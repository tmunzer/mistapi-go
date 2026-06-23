
# Org Setting Celona

Integration settings for Celona

## Structure

`OrgSettingCelona`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiKey` | `*string` | Optional | Credential used by Mist for the Celona integration |
| `ApiPrefix` | `*string` | Optional | Celona API prefix configured for the integration |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingCelona := models.OrgSettingCelona{
        ApiKey:               models.ToPointer("$2a$04$OkaLCoJn6rDjR8ha.oduQVDST3.kJNIrte"),
        ApiPrefix:            models.ToPointer("cc3273fcb016470e"),
    }

}
```

