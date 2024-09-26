
# Admin Privilege View Enum

Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.

You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.

Below are the list of supported UI views. Note that this is UI only feature
Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.

You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.

Below are the list of supported UI views. Note that this is UI only feature

| UI View | Required Role | Description |
| --- | --- | --- |
| `reporting` | `read` | full access to all analytics tools |
| `marketing` | `read` | can view analytics and location maps |
| `super_observer` | `read` | can view all the organization except the subscription page |
| `location` | `write` | can view and manage location maps, can view analytics |
| `security` | `write` | can view and manage site labels, policies and security |
| `switch_admin` | `helpdesk` | can view and manage Switch ports, can view wired clients |
| `mxedge_admin` | `admin` | can view and manage Mist edges and Mist tunnels |
| `lobby_admin` | `admin` | full access to Org and Site Pre-shared keys |

## Enumeration

`AdminPrivilegeViewEnum`

## Fields

| Name |
|  --- |
| `lobby_admin` |
| `location` |
| `marketing` |
| `mxedge_admin` |
| `reporting` |
| `security` |
| `super_observer` |
| `switch_admin` |

