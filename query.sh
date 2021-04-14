curl -g \
-X POST \
-H "Content-Type: application/json" \
-d '{"query":"query{people {name description}}"}' \
http://localhost:8080/query

echo 


# query {
#	links{
#    title
#    address,
#    user{
#      name
#    }
#  }
# }

# query {
#	people{
#    name
#    description
#  }
# }
