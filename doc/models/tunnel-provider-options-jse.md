
# Tunnel Provider Options Jse

For jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsJse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumUsers` | `*int` | Optional | - |
| `OrgName` | `*string` | Optional | JSE Organization name. The list of available organizations can be retrieved with the [Get Org JSE Info](../../doc/controllers/orgs-jse.md#get-org-jse-info) API Call |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_users": 5,
  "org_name": "JSE_ORG1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

