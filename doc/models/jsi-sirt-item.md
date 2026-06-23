
# Jsi Sirt Item

Juniper Security Intelligence SIRT advisory item

## Structure

`JsiSirtItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CvssScore` | `*float64` | Optional | Common Vulnerability Scoring System score for the SIRT advisory |
| `Id` | `*string` | Optional | Unique SIRT or JSA advisory identifier from Juniper Support Insights |
| `Models` | `[]string` | Optional | Device models affected by the SIRT advisory |
| `Problem` | `*string` | Optional | Issue details described by the SIRT advisory |
| `PublishedDate` | `*int` | Optional | Release date of the SIRT issue |
| `ReleaseNotes` | `*string` | Optional | Release notes if any |
| `Severity` | `*string` | Optional | Security severity assigned to the SIRT advisory |
| `Solution` | `*string` | Optional | Recommended fix or remediation for the security issue |
| `Title` | `*string` | Optional | Summary title for the SIRT advisory |
| `UpdatedDate` | `*int` | Optional | Time when the JSA advisory was last updated |
| `Versions` | `[]string` | Optional | Software versions affected by the SIRT advisory |
| `Workaround` | `*string` | Optional | Mitigation or workaround guidance for the SIRT advisory |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiSirtItem := models.JsiSirtItem{
        CvssScore:            models.ToPointer(float64(41.8)),
        Id:                   models.ToPointer("JSA100053"),
        Models:               []string{
            "models8",
            "models9",
        },
        Problem:              models.ToPointer("problem2"),
        PublishedDate:        models.ToPointer(220),
    }

}
```

