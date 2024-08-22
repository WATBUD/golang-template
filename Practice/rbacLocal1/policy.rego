package rbac

# Importing the future keywords package to use the `in` keyword for iteration
import future.keywords.in

# Default deny policy
default allow = false

# Defining roles and their associated permissions
roles = {
  "admin": {
    # Admin has permissions to create and delete channels, and manage users
    "permissions": ["create_channel", "delete_channel", "manage_users"]
  },
  "moderator": {
    # Moderator has permissions to create and delete channels
    "permissions": ["create_channel", "delete_channel"]
  },
  "user": {
    # User has permissions to join channels and send messages
    "permissions": ["join_channel", "send_message"]
  }
}

# Mapping users to their roles
users = {
  "alice": ["admin"],      # Alice is an admin
  "bob": ["moderator"],    # Bob is a moderator
  "eve": ["user"]          # Eve is a regular user
}

# Allow rule to check if the user has the required permission for the action
allow {
  # Check if the user's role is in the defined users' roles
  some role in users[input.user]
  # Check if the permission for the action exists in the role's permissions
  some perm in roles[role].permissions
  # Ensure the permission matches the action being requested
  perm == input.action
}
``
