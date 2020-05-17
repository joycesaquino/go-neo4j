package dao

// TODO Fazer um query builder passando somente o objeto
const (
	InsertServiceWithUserRelation = "CREATE (service:Service { " +
		"id: $id, description: $description, value: $value, initialDateTime: $initialDateTime, finalDateTime: $finalDateTime, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions,createdAt: $createdAt}" +
		")-[r:PERTENCE]->(user:User {name: $name, email: $email})" +
		" RETURN service.id, service.description,r"

	FindServiceById = "MATCH (service:Service {id: $id}) RETURN service"
)
