all: build-app
.PHONY: build-app
build-app: app.py requirements.txt templates/index.html Dockerfile
	 docker build -t unknowntpo/flask-pic .
	 docker container stop flask-pic
	 docker container rm flask-pic
	 docker run -d -p 8888:5000 --name "flask-pic" unknowntpo/flask-pic
	 

