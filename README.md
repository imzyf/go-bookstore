# go-bookstore

## CURL

```bash
curl -X POST -H "Content-Type:application/json" -d '{"id":"978-7-115-31897-8","name":"图解 TCP/IP","authors":["竹下隆史","村山公保"],"press":"人民邮电出版社"}' localhost:8888/book

curl -X GET -H "Content-Type:application/json" localhost:8888/book/978-7-115-31897-8

curl -X PUT -H "Content-Type:application/json" -d '{"id":"978-7-115-31897-8","name":"图解 TCP/IP v2"}' localhost:8888/book/978-7-115-31897-8

curl -X POST -H "Content-Type:application/json" -d '{"id":"978-7-115-31897-9","name":"图解 TCP/IP v10","authors":["竹下隆史","村山公保"],"press":"人民邮电出版社"}' localhost:8888/book

curl -X GET -H "Content-Type:application/json" localhost:8888/books

curl -X DELETE -H "Content-Type:application/json" localhost:8888/book/978-7-115-31897-8
```