
# Response Org Search Item

Organization record returned by MSP organization search

## Structure

`ResponseOrgSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `Name` | `*string` | Optional, Read-only | Display name of the organization |
| `NumAps` | `*int` | Optional, Read-only | Number of APs in the organization |
| `NumGateways` | `*int` | Optional, Read-only | Number of gateways in the organization |
| `NumSites` | `*int` | Optional, Read-only | Number of sites in the organization |
| `NumSwitches` | `*int` | Optional, Read-only | Number of switches in the organization |
| `NumUnassignedAps` | `*int` | Optional, Read-only | Number of APs in organization inventory that are not assigned to a site |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SubAnaEntitled` | `*int` | Optional, Read-only | Number of SUB-ANA subscriptions entitled for the organization |
| `SubAnaRequired` | `*int` | Optional, Read-only | Number of SUB-ANA subscriptions required by the organization |
| `SubAstEntitled` | `*int` | Optional, Read-only | Number of SUB-AST subscriptions entitled for the organization |
| `SubAstRequired` | `*int` | Optional, Read-only | Number of SUB-AST subscriptions required by the organization |
| `SubEngEntitled` | `*int` | Optional, Read-only | Number of SUB-ENG subscriptions entitled for the organization |
| `SubEngRequired` | `*int` | Optional, Read-only | Number of SUB-ENG subscriptions required by the organization |
| `SubEx12Required` | `*int` | Optional, Read-only | Number of SUB-EX12 subscriptions required by the organization |
| `SubInsufficient` | `*bool` | Optional, Read-only | If this org has sufficient subscription |
| `SubManEntitled` | `*int` | Optional, Read-only | Number of SUB-MAN subscriptions entitled for the organization |
| `SubManRequired` | `*int` | Optional, Read-only | Number of SUB-MAN subscriptions required by the organization |
| `SubMeEntitled` | `*int` | Optional, Read-only | Number of SUB-ME subscriptions entitled for the organization |
| `SubVnaEntitled` | `*int` | Optional, Read-only | Number of SUB-VNA subscriptions entitled for the organization |
| `SubVnaRequired` | `*int` | Optional, Read-only | Number of SUB-VNA subscriptions required by the organization |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `TrialEnabled` | `*bool` | Optional, Read-only | If this org is under trial period |
| `UsageTypes` | `[]string` | Optional, Read-only | Subscription usage type codes enabled for an organization |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseOrgSearchItem := models.ResponseOrgSearchItem{
        MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        Name:                 models.ToPointer("name0"),
        NumAps:               models.ToPointer(108),
        NumGateways:          models.ToPointer(144),
        NumSites:             models.ToPointer(82),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

