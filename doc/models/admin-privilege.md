
# Admin Privilege

Privilieges settings

## Structure

`AdminPrivilege`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MspLogoUrl` | `*string` | Optional | logo of the MSP (if the MSP belongs to an Advanced tier) |
| `MspName` | `models.Optional[string]` | Optional | name of the MSP (if the org belongs to an MSP) |
| `MspUrl` | `*string` | Optional | custom url of the MSP (if the MSP belongs to an Advanced tier) |
| `Name` | `*string` | Optional | name of the org/site/MSP depending on object scope |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `OrgName` | `*string` | Optional | name of the org (for a site belonging to org) |
| `OrggroupIds` | `[]uuid.UUID` | Optional | if `scope`==`orggroup` |
| `Role` | [`models.AdminPrivilegeRoleEnum`](../../doc/models/admin-privilege-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.AdminPrivilegeScopeEnum`](../../doc/models/admin-privilege-scope-enum.md) | Required | enum: `msp`, `org`, `orggroup`, `site`, `sitegroup` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SitegroupIds` | `[]uuid.UUID` | Optional | - |
| `Views` | [`*models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br>Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |

## Example (as JSON)

```json
{
  "msp_id": "0a9c626d-be6d-4e00-8402-b11370c5a308",
  "org_id": "6d03c926-72e3-40cd-92ab-4560c1f4e33d",
  "role": "installer",
  "scope": "site",
  "site_id": "52b50564-8821-4c3e-97be-5061c7760002",
  "msp_logo_url": "msp_logo_url2",
  "msp_name": "msp_name8",
  "msp_url": "msp_url4",
  "name": "name4"
}
```

