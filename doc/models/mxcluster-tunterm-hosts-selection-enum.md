
# Mxcluster Tunterm Hosts Selection Enum

Ordering of tunterm_hosts for mxedge within the same mxcluster. enum:

* `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC
* `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)
* `ordered`: order decided by tunterm_hosts_order

## Enumeration

`MxclusterTuntermHostsSelectionEnum`

## Fields

| Name |
|  --- |
| `ordered` |
| `shuffle` |
| `shuffle-by-site` |

