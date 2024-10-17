# Lab

## Errors

❯ terraform plan
╷
│ Error: Unsupported attribute
│
│ on main.tf line 8, in resource "decort_resgroup" "lab":
│ 8: gid = data.decort_locations_list.ll.gid
│
│ This object has no argument, nested block, or exported attribute named "gid".
│ Did you mean "id"?
