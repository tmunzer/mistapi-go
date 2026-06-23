
# Marvis Config Feedback Response

Response after submitting feedback on a Marvis config action

## Structure

`MarvisConfigFeedbackResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FeedbackNote` | `*string` | Optional | The note provided with the feedback |
| `FeedbackType` | `*string` | Optional | The feedback type that was submitted |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisConfigFeedbackResponse := models.MarvisConfigFeedbackResponse{
        FeedbackNote:         models.ToPointer("feedback_note4"),
        FeedbackType:         models.ToPointer("feedback_type2"),
    }

}
```

