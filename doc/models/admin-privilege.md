
# Admin Privilege

Privilieges settings

*This model accepts additional fields of type interface{}.*

## Structure

`AdminPrivilege`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MspId` | `*uuid.UUID` | Optional | Required if `scope`==`msp` |
| `MspLogoUrl` | `*string` | Optional | Logo of the MSP (if the MSP belongs to an Advanced tier) |
| `MspName` | `models.Optional[string]` | Optional | Name of the MSP (if the org belongs to an MSP) |
| `MspUrl` | `*string` | Optional | Custom url of the MSP (if the MSP belongs to an Advanced tier) |
| `Name` | `*string` | Optional | Name of the org/site/MSP depending on object scope |
| `OrgId` | `*uuid.UUID` | Optional | Required if `scope`==`org` |
| `OrgName` | `*string` | Optional | Name of the org (for a site belonging to org) |
| `OrggroupIds` | `[]uuid.UUID` | Optional | If `scope`==`orggroup` |
| `Role` | [`models.AdminPrivilegeRoleEnum`](../../doc/models/admin-privilege-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.AdminPrivilegeScopeEnum`](../../doc/models/admin-privilege-scope-enum.md) | Required | enum: `msp`, `org`, `orggroup`, `site`, `sitegroup` |
| `SiteId` | `*uuid.UUID` | Optional | Required if `scope`==`site` |
| `SitegroupIds` | `[]uuid.UUID` | Optional | - |
| `Views` | [`[]models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  <br>You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  <br>Below are the list of supported UI views. Note that this is UI only feature.<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "role": "installer",
  "scope": "site",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "msp_logo_url": "msp_logo_url2",
  "msp_name": "msp_name8",
  "msp_url": "msp_url4",
  "name": "name4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

