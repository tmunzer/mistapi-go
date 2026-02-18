
# Jsi Sirt Item

SIRT advisory item

## Structure

`JsiSirtItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CvssScore` | `*float64` | Optional | CVSS score |
| `Id` | `*string` | Optional | ID of the SIRT |
| `Models` | `[]string` | Optional | OS models affected |
| `Problem` | `*string` | Optional | Problem description |
| `PublishedDate` | `*int` | Optional | Release date of the SIRT issue |
| `ReleaseNotes` | `*string` | Optional | Release notes if any |
| `Severity` | `*string` | Optional | Severity of the issue |
| `Solution` | `*string` | Optional | Solution for the security issue |
| `Title` | `*string` | Optional | Title of the SIRT |
| `UpdatedDate` | `*int` | Optional | JSA updated timestamp |
| `Versions` | `[]string` | Optional | OS versions affected |
| `Workaround` | `*string` | Optional | Workaround provided |

## Example (as JSON)

```json
{
  "id": "JSA100053",
  "cvss_score": 193.42,
  "models": [
    "models0",
    "models1",
    "models2"
  ],
  "problem": "problem4",
  "published_date": 162
}
```

