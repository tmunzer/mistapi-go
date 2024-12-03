
# License Usage Site

*This model accepts additional fields of type interface{}.*

## Structure

`LicenseUsageSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgEntitled` | `map[string]int` | Required | license entitlement for the entire org |
| `SvnaEnabled` | `bool` | Required | eligibility for the Switch SLE |
| `TrialEnabled` | `bool` | Required | - |
| `Usages` | `map[string]int` | Required | subscriptions and their quantities |
| `VnaEligible` | `bool` | Required | eligibility for the AP/Client SLE |
| `VnaUi` | `bool` | Required | if True, Conversational Assistant and Marvis Action available |
| `WvnaEligible` | `bool` | Required | eligibility for the WAN SLE |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_entitled": {
    "SUB-LOC": 30,
    "SUB-MAN": 60
  },
  "svna_enabled": false,
  "trial_enabled": false,
  "usages": {
    "SUB-LOC": 30,
    "SUB-MAN": 60
  },
  "vna_eligible": false,
  "vna_ui": false,
  "wvna_eligible": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

