CONTAINER_NAME=sqlbenchmark

all: benchmark

benchmark: clean
	docker run --name $(CONTAINER_NAME) -e MYSQL_ALLOW_EMPTY_PASSWORD=TRUE -e MYSQL_DATABASE="benchmark" -p 3306:3306 mysql

container:
	docker exec -it $(CONTAINER_NAME) bash -c "mysql -u root"

server: src/server.go
	cd src && go run server.go

plot: plot.gp
	gnuplot plot.gp

.PHONY: clean

clean:
	docker stop $(CONTAINER_NAME) || true
	docker rm -f $(CONTAINER_NAME)