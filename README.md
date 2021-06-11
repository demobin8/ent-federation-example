### ent federation example

### 1. start role and user servie

`cd user`
`go run server.go`

`cd role`
`go run server.go`

### 2. start gateway

`yarn start`

### 3. query

address:
http://127.0.0.1:8080/

mutation
```
# Write your query or mutation here

# mutation {
#   createRolePermission(rolePermission :{
#     roleId: 1
#     permissionId: 6
#   })
# }

# query {
#   roles(orderBy: {
#     direction: ASC
#     field: NAME
#   }) {
#     totalCount
#     pageInfo {
#       hasNextPage
#       hasPreviousPage
#       endCursor
#       startCursor
#     }
#     edges {
#       node {
#         id
#         name
#         remark
#       }
#       cursor
#     }
#   }
# }

# query {
#   user {
#     id
#     name
#   }
# }
      
query {
	userRoleDetail (userId: 1) {
		id
    name
    roles {
      role {
        id
        name
        remark
      }
      permissions {
        id
        name
        remark
        code
      }
    }
  }
}
```
