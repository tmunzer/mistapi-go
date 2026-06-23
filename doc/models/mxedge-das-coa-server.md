
# Mxedge Das Coa Server

CoA or Disconnect-Message client allowed to contact Mist Edge DAS

## Structure

`MxedgeDasCoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | Whether to disable Event-Timestamp Check<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether this DAS CoA or Disconnect-Message client is enabled |
| `Host` | `*string` | Optional | Server host allowed to send CoA or Disconnect-Message requests to Mist Edges |
| `Port` | `*int` | Optional | UDP port where Mist Edges accept CoA or Disconnect-Message requests from this host<br><br>**Default**: `3799` |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Secret` | `*string` | Optional | Shared secret used by this DAS CoA or Disconnect-Message client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeDasCoaServer := models.MxedgeDasCoaServer{
        DisableEventTimestampCheck:  models.ToPointer(false),
        Enabled:                     models.ToPointer(false),
        Host:                        models.ToPointer("host6"),
        Port:                        models.ToPointer(3799),
        RequireMessageAuthenticator: models.ToPointer(false),
    }

}
```

