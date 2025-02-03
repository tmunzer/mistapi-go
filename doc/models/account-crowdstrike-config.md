
# Account Crowdstrike Config

OAuth linked CrowdStrike apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountCrowdstrikeConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client ID |
| `ClientSecret` | `string` | Required | Customer account api client Secret |
| `CustomerId` | `string` | Required | Customer id of an admin |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_id": "client_id4",
  "client_secret": "client_secret0",
  "customer_id": "customer_id0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

