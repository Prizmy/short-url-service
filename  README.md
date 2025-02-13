go run main.go -db=postgres
go run main.go -db=memory

docker-compose build --no-cache 
docker-compose up  
