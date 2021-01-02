# Microservices using golang beego framework

#### Build
```
docker build -t app  .
```

#### Docker Run Command
```
   docker run -d -p 8080:8080 --name app app

```

#### Swagger
`http://localhost:8080/swagger`

#### The master node should be tagged with dedicated: master

#### Directory structure 
Controller : contain API definitions 
Models: contain business logic 
Constants : container file paths 
Routers: contain API http route 
Scripts : contain centos and ubuntu preparation 

