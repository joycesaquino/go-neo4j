# Dera Services API

API de serviços que são oferecidos pelos usuários do app.

http.PostMethod /service

#### Jwt
 
     payload : {
       "user": "username",
       "email": "useremail@gmail.com"
      }
 
 
     HMACSHA256(
       base64UrlEncode(header) + "." +
       base64UrlEncode(payload),
       secret-api
     )
 
     header: {
       "alg": "HS256",
       "typ": "JWT"
     }

#### Body 
     { 
        "id": auto-generated,
        "description": string,
        "value": float,
        "initialDateTime": neo4j.LocalDateTime, 
        "finalDateTime": neo4j.LocalDateTime,
        "maxSubscriptions": int,
        "minSubscriptions": int,
        "createdAt": neo4j.LocalDateTime
     }

#### Response

     201 created
     400 Bad Request
     401 Unauthorized

 