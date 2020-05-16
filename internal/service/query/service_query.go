package query

// TODO Fazer um query builder passando somente o objeto
const InsertServiceQuery = "CREATE (service:Service { " +
	"id: $id, description: $description, value: $value, initialDateTime: $initialDateTime, finalDateTime: $finalDateTime, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions,createdAt: $createdAt}" +
	")-[r:PERTENCE]->(user:User {name: $name, email: $email})" +
	" RETURN service.id, service.description,r"
