
# Response Config History Search Item Radio

Radio config history detail

## Structure

`ResponseConfigHistorySearchItemRadio`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `string` | Required | Radio band for this config history detail |
| `Channel` | `int` | Required | Configured channel for this radio |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseConfigHistorySearchItemRadio := models.ResponseConfigHistorySearchItemRadio{
        Band:                 "band6",
        Channel:              80,
    }

}
```

