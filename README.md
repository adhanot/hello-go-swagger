# hello-go-swagger

Hello-go-swagger is a Golang Template to build fast RESTful API by code generation. The complete stack creates a powerful and secure framework.

It use go-swagger (https://goswagger.io/), ENT Framework (https://entgo.io/) and JWT (https://github.com/dgrijalva/jwt-go).

Hello-go-swagger generate the client and the server. 

If you want modify the API go-swagger server, you need to modify the : 
```bash
/gen/api/swagger.yaml
```

If you want add ENT database schema, you can add file in : 
```bash
/gen/api/ent
```


## Installation

Use need to use docker-compose command :

```bash
docker-compose build --no-cache  && docker-compose up
```

## Usage

Hello-go-swagger launch 3 API routes.
The routes are :
```bash
http://0.0.0.0:8081/hello
http://0.0.0.0:8081/register
http://0.0.0.0:8081/login
```
You have a complete register and login authentification with JWT token and you have a "/home" route. You can only access on /home route with your JWT token. You need to be log before access to /home.

After modify the swagger.yaml or add new database schema, you just need go in the "scripts" folder and launch :
```bash
sh generate.sh
```
