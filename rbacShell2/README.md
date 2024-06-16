
# 在MongoDB Shell中創建數據
# use rbac_db
# db.roles.insertMany([
#   {
#     "role": "admin",
#     "permissions": ["create_channel", "delete_channel", "manage_users"]
#   },
#   {
#     "role": "moderator",
#     "permissions": ["create_channel", "delete_channel"]
#   },
#   {
#     "role": "user",
#     "permissions": ["join_channel", "send_message"]
#   }
# ])
# 
# db.users.insertMany([
#   {
#     "user": "alice",
#     "roles": ["admin"]
#   },
#   {
#     "user": "bob",
#     "roles": ["moderator"]
#   },
#   {
#     "user": "eve",
#     "roles": ["user"]
#   }
# ])
