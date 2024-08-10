
# Privilege Org

Privilieges settings

## Structure

`PrivilegeOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Role` | [`models.PrivilegeOrgRoleEnum`](../../doc/models/privilege-org-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeOrgScopeEnum`](../../doc/models/privilege-org-scope-enum.md) | Required | enum: `org`, `site`, `sitegroup` |
| `SiteId` | `*uuid.UUID` | Optional | if `scope`==`site` |
| `SitegroupId` | `*uuid.UUID` | Optional | if `scope`==`sitegroup` |
| `Views` | [`*models.PrivilegeOrgViewsEnum`](../../doc/models/privilege-org-views-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br>Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br><br>\| UI View \| Description \|<br>\| --- \| --- \|<br>\| `reporting` \| full access to all analytics tools \|<br>\| `marketing` \| can view analytics and location maps \|<br>\| `location` \| can view and manage location maps \|<br>\| `security` \| can view and manage WLAN, rogues and authentication \|<br>\| `switch_admin` \| can view and manage Switch ports \|<br>\| `mxedge_admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| full access to Org and Site Pre-shared keys \| |

## Example (as JSON)

```json
{
  "role": "write",
  "scope": "sitegroup",
  "site_id": "0000169c-0000-0000-0000-000000000000",
  "sitegroup_id": "00002114-0000-0000-0000-000000000000",
  "views": "marketing"
}
```

