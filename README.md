##### To run container

```console
sudo docker-compose exec -d

sudo docker-compose exec postgres sh

> hostname -i
```

---

##### To execute project GO in other path

```console
go mod init go-rest-api

or

go mod init github.com/LucasSantosDev/go-rest-api
```

---

##### To build and execute

```console
go run main.go
```

---

##### To test list personalities

```console
curl -H "Content-Type: application/json" -X GET http://localhost:8000/api/personalities
```

---

##### To test show specific personalitie

```console
curl -H "Content-Type: application/json" -X GET http://localhost:8000/api/personalities/1
```

---

##### To test create personalitie

```console
curl -H "Content-Type: application/json" -X POST -d '{"name":"Joana","history":"Outra história"}' http://localhost:8000/api/personalities
```

---

##### To test delete personalitie

```console
curl -H "Content-Type: application/json" -X DELETE http://localhost:8000/api/personalities/4
```

---

##### To test update personalitie

```console
curl -H "Content-Type: application/json" -X PUT -d '{"name":"Noem atualizado","history":"História atualizada"}' http://localhost:8000/api/personalities/2
```

---