curl -g \
-X POST \
-H "Content-Type: application/json" \
-d '{"query":"mutation {createPerson(input: {name: \"Blahmy Bla\", description: \"a new person\"}) {name, description}}"}' \
http://localhost:8080/query

echo 


# mutation {
# createLink(input: {title: "new link", address:"http://address.org"}){
#    title,
#    user{
#      name
#    }
#    address
#  }
#}

# mutation {
# createPerson(input: {name: "Blahmy Bla", description: "a new person"}){
#    name,
#    description
#  }
#}