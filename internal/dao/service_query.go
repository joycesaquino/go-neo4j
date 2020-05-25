package dao

// TODO Fazer um query builder passando somente o objeto

//Relations
const (
	provide   = "PROVIDE"
	subscribe = "SUBSCRIBE"
	belongs   = "BELONGS"
)

const (
	CreateService = "CREATE (n:service { " +
		"id: $id, description: $description, value: $value,createdAt: $createdAt}" +
		")-[r:" + provide + "]->(user:User {name: $name, email: $email})" +
		" RETURN n.id, n.description,r"

	InsertClassWithServiceRelation = "CREATE (n:class { " +
		"id: $id, description: $description, value: $value, initialDateTime: $initialDateTime, finalDateTime: $finalDateTime, subscriptions: $subscriptions,minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions,createdAt: $createdAt}" +
		")-[r:" + belongs + "]->(n:service {id: $id})" +
		" RETURN n.id, n.description,r"

	FindServiceById = "MATCH (n:class {id: $id}) RETURN n"

	UserSubscribeService = "MATCH (n:class {id: $id}) SET n. RETURN n"
)
