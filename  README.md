В docker-compose можно изменять парметр DB_TYPE=memory или на postgres для изминения типа хранилища.

run main.go -db=postgres
go run main.go -db=memory

docker-compose build
docker-compose up  
