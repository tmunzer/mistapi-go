
# Mxcluster Tunterm Hosts Selection Enum

Ordering of tunterm_hosts for mxedge within the same mxcluster.

* When `shuffle`, the ordering of tunterm_hosts is randomized by the device’s MAC.
* When `shuffle-by-site`, we shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels).
* When `ordered`, the order is decided by tunterm_hosts_order

## Enumeration

`MxclusterTuntermHostsSelectionEnum`

## Fields

| Name |
|  --- |
| `shuffle` |
| `shuffle-by-site` |
| `ordered` |

