
# Privilege Org

Privileges settings

## Structure

`PrivilegeOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | If `scope`==`org` |
| `Role` | [`models.PrivilegeOrgRoleEnum`](../../doc/models/privilege-org-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeOrgScopeEnum`](../../doc/models/privilege-org-scope-enum.md) | Required | enum: `org`, `site`, `sitegroup`, `orgsites` |
| `SiteId` | `*uuid.UUID` | Optional | If `scope`==`site` |
| `SitegroupId` | `*uuid.UUID` | Optional | If `scope`==`sitegroup` |
| `View` | `*string` | Optional | Used for backward compatibility. Use `views` instead. |
| `Views` | [`[]models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  <br>You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  <br>Below are the list of supported UI views. Note that this is UI only feature.<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |

## Example (as JSON)

```json
{
  "org_id": "0000270e-0000-0000-0000-000000000000",
  "role": "write",
  "scope": "org",
  "site_id": "0000169c-0000-0000-0000-000000000000",
  "sitegroup_id": "00002114-0000-0000-0000-000000000000",
  "view": "view0",
  "views": [
    "location"
  ]
}
```

