package rbac

import future.keywords.in

default allow = false

allow {
  some role in input.roles[input.user]
  some perm in input.permissions[role]
  perm == input.action
}
