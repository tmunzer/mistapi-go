
# Response Org Search Item

## Structure

`ResponseOrgSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | org name |
| `NumAps` | `*int` | Optional | - |
| `NumGateways` | `*int` | Optional | - |
| `NumSites` | `*int` | Optional | - |
| `NumSwitches` | `*int` | Optional | - |
| `NumUnassignedAps` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SubAnaEntitled` | `*int` | Optional | - |
| `SubAnaRequired` | `*int` | Optional | - |
| `SubAstEntitled` | `*int` | Optional | - |
| `SubAstRequired` | `*int` | Optional | - |
| `SubEngEntitled` | `*int` | Optional | - |
| `SubEngRequired` | `*int` | Optional | - |
| `SubEx12Required` | `*int` | Optional | - |
| `SubInsufficient` | `*bool` | Optional | if this org has sufficient subscription |
| `SubManEntitled` | `*int` | Optional | - |
| `SubManRequired` | `*int` | Optional | - |
| `SubMeEntitled` | `*int` | Optional | - |
| `SubVnaEntitled` | `*int` | Optional | - |
| `SubVnaRequired` | `*int` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `TrialEnabled` | `*bool` | Optional | if this org is under trial period |
| `UsageTypes` | `[]string` | Optional | a list of types that enabled by usage |

## Example (as JSON)

```json
{
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "name": "name6",
  "num_aps": 126,
  "num_gateways": 126,
  "num_sites": 64
}
```

