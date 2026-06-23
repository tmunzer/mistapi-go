
# Response Assign Success

Assignment operation success response

## Structure

`ResponseAssignSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Success` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAssignSuccess := models.ResponseAssignSuccess{
        Success:              []string{
            "success4",
        },
    }

}
```

