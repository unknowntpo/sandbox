all: build

build:
	sudo docker build -t testing_restarts ./

run:
	sudo docker run -d --name testing_restarts testing_restarts

# Run the container with restart policy
run/restart: 
	sudo docker run -d --name testing_restarts --restart always testing_restarts
