# Binary tree
This is a binary tree implemented in Golang. 
All the public interface is exposed through a web api using [Gorilla Mux](https://github.com/gorilla/mux)

## Building and running
The application runs a web server hosted on port `8080` by default.
To build the application, run the following script in the root folder
```bash
make build
```

To run the built application just run the executable file with
```bash
bin/binary-tree
```

To access it you can either call the endpoints using Postman or curl:

### Getting Health endpoint
```bash
curl -X GET -I http://localhost:8080/health
```

### Inserting a value
```bash
curl -X POST -I http://localhost:8080/bt/insert/5
```

### Checking if the Binary Tree contains a value
```bash
curl -X GET http://localhost:8080/bt/contains/5
```

### Printintg the Binary Tree
```bash
curl -X GET http://localhost:8080/bt
```

## Testing
Run the following script in the root folder
```bash
go test -v -cover ./...
```
or 
```bash
make test
```