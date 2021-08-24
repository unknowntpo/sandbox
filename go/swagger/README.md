An experiment of using go-swagger and swagger UI docker container

1. Follow the documentation[1], and write the spec.

2. run following command:

swagger generate spec -o ./swagger.json

to generate swagger.json.

3. Run swagger UI container to serve the spec.

$ docker-compose up -d

4. Open up http://localhost:8080/swagger/ to read the documentation, the port and endpoint
can be changed inside docker-compose.yml.

Ref:
    - [1](https://goswagger.io/generate/spec.html)
