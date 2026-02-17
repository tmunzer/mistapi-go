
# Jsi Pbn Item

PBN (Problem Bug Notification) advisory item

## Structure

`JsiPbnItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BugType` | `*string` | Optional | Type of the bug (Day-1, Regression) |
| `CustomerRisk` | `*string` | Optional | Risk level |
| `FixedIn` | `*string` | Optional | Release in which the issue was fixed |
| `Id` | `*string` | Optional | ID of the PBN |
| `IntroducedIn` | `*string` | Optional | Release introduced in |
| `Models` | `[]string` | Optional | OS models affected |
| `ProductFamily` | `*string` | Optional | Product family affected |
| `ReleaseNotes` | `*string` | Optional | Release notes for this PBN |
| `Restoration` | `*string` | Optional | Restoration steps |
| `Title` | `*string` | Optional | Title of the issue |
| `UpdatedDate` | `*int` | Optional | PBN updated timestamp |
| `Versions` | `[]string` | Optional | OS versions affected |
| `Workaround` | `*string` | Optional | Workaround for this issue |
| `WorkaroundProvided` | `*string` | Optional | Any workaround available |

## Example (as JSON)

```json
{
  "id": "1403338",
  "bug_type": "bug_type0",
  "customer_risk": "customer_risk4",
  "fixed_in": "fixed_in8",
  "introduced_in": "introduced_in2"
}
```

