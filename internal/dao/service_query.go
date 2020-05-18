package dao

// TODO Fazer um query builder passando somente o objeto

//Relations
const (
	UserProvideService   = "PROVIDE"
	UserSubscribeService = "SUBSCRIBE"
)

const (
	InsertServiceWithUserRelation = "CREATE (n:service { " +
		"id: $id, description: $description, value: $value, initialDateTime: $initialDateTime, finalDateTime: $finalDateTime, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions,createdAt: $createdAt}" +
		")-[r:" + UserProvideService + "]->(user:User {name: $name, email: $email})" +
		" RETURN n.id, n.description,r"

	FindServiceById = "MATCH (n:service {id: $id}) RETURN n"
)
