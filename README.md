# GEORGIA-HELLO-WORLD

## **required**: 

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/get-started/)
--------

### Endpoints

##### /
Methods: GET

a simple hello world message

returns ```{"message":"Hello World!"}``` response 

---------

#### /status
Methods: GET

metadata about the running application. This is "baked in" at build.

returns ```{"my-application":{"version":"1.0.0","description":"basic hello-world app in development environment","sha":"a0ce191"}}``` 


### build and run locally

can ommit build args if testing locally

```docker build --build-arg version=<version> --build-arg sha="<commit> --build-arg description=<description> -t georgia-hello-world:<test-tag>  .```

run:


```docker run --rm -p 80:8001 georgia-hello-world```

----------

## build and publish an image for testing (non CI release)

auth with ghcr registry. e.g with PAK:

``` echo $GHCR_TOKEN | docker login ghcr.io -u <user-name> --password-stdin  ```

run build and publish script

```./bin/build_and_publish.sh <environment> <image-tag>```


>NOTE: if on macOS with newer architecture you may need to explicitely set platform:

``` --platform linux/amd64 ```

Testing

```./bin/test.sh``` or ```go test -v ./...```
