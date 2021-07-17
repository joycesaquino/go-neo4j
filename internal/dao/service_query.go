package dao

// TODO Fazer um query builder passando somente o objeto

//Relations
const (
	provide   = "PROVIDE"
	belongs   = "BELONGS"
	subscribe = "SUBSCRIBE"
)

const (
	InsertService = "CREATE (service:service { " +
		"id: $id, " +
		"description: $description, " +
		"value: $value," +
		"createdAt: $createdAt}" +
		")-[r:" + provide + "]->(user:User {name: $name, email: $email})" +
		" RETURN service.id, service.description,r"

	InsertClassWithServiceRelation = "CREATE (class:class { " +
		"id: $id, " +
		"description: $description, " +
		"value: $value, " +
		"initialDateTime: $initialDateTime, " +
		"finalDateTime: $finalDateTime, " +
		"subscriptions: $subscriptions, " +
		"minSubscriptions: $minSubscriptions, " +
		"maxSubscriptions: $maxSubscriptions," +
		"createdAt: $createdAt}" +
		")-[r:" + belongs + "]->(service:service {id: $id})" +
		" RETURN class.id, class.description,r"

	FindServiceById = "MATCH (class:class {id: $id}) RETURN class"

	UserSubscribeService = "MATCH (n:class {id: $id}) SET n. RETURN n"
)
