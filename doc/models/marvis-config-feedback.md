
# Marvis Config Feedback

Feedback submission for a Marvis config action

## Structure

`MarvisConfigFeedback`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | `*string` | Optional | Free-text note about the feedback |
| `Type` | [`*models.MarvisConfigFeedbackTypeEnum`](../../doc/models/marvis-config-feedback-type-enum.md) | Optional | Feedback type. enum: `invalid` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisConfigFeedback := models.MarvisConfigFeedback{
        Note:                 models.ToPointer("note2"),
        Type:                 models.ToPointer(models.MarvisConfigFeedbackTypeEnum_INVALID),
    }

}
```

