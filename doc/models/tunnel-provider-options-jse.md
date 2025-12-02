
# Tunnel Provider Options Jse

For jse-ipsec, this allows provisioning of adequate resource on JSE. Make sure adequate licenses are added

## Structure

`TunnelProviderOptionsJse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumUsers` | `*int` | Optional | - |
| `OrgName` | `*string` | Optional | JSE Organization name. The list of available organizations can be retrieved with the [Get Org JSE Info]($e/Orgs%20JSE/getOrgJseInfo) API Call |

## Example (as JSON)

```json
{
  "num_users": 5,
  "org_name": "JSE_ORG1"
}
```

