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


#### /status
Methods: GET

metadata about the running application. This is "baked in" at build.

returns ```{"my-application":{"version":"1.0.0","description":"basic hello-world app in development environment","sha":"a0ce191"}}``` 

-------
### build and run locally

can ommit build args if testing locally

```docker build --build-arg version=<version> --build-arg sha=<commit> --build-arg description=<description> -t georgia-hello-world:<test-tag>  .```

run:


```docker run --rm -p 80:8001 georgia-hello-world```

----------

## publish new release
- update the VERSION
- merge PR into main
- tag main with release version:

```git tag vx.x.x```
- push tag:

```git push origin vx.x.x```

- create a release in github UI and select latest tag -> pipeline will run the build and release


## build and publish an image for testing (non CI release)

auth with ghcr registry. e.g with PAK:

``` echo $GHCR_TOKEN | docker login ghcr.io -u <user-name> --password-stdin  ```

run build and publish script

```./bin/build_and_publish.sh <environment> <image-tag>```


>NOTE: if on macOS with newer architecture you may need to explicitely set platform:

``` --platform linux/amd64 ```

Testing

```./bin/test.sh``` or ```go test -v ./...```


Limitations

### CI

the build pipeline will run test and linting steps on every branch, this is good to validate work in process. The build and publish steps however run on *every merge to main* - which might be fine, but if there are changes to the README.md or minor changes to improve the code that doesn't mean a patch is required it's a waste.  

perhaps I could use tags on releases and differentiate on this basis.

I don't like how I read the description in the bin/build_and_publish.sh
this is probably fine for a small project - but for something bigger and with more configuration it would be worth utilizing some templating tooling to read various env configurations across environments.  

CI is very coupled with github actions. I have tried to put most of the "work" in scripts because I think it's nice to be able to run them locally. 

### the code
I am not a go programmer, so while there are issues I am aware of there are likely many I am not.

error handling is not ideal.

One thing I'm not enthusiastic about is I have included the [gorrila/mux](https://github.com/gorilla/mux) pkg because I wanted a quick way to enforce what methods are allowed at the 'top' routing level. I notice they are looking for a new maintainer and think it's generally best not to rely on external tools for simple things like this, and I'm not sure it gives me a lot, but this was faster for me because I'm not a go-wiz.

### testing
 I like to rely and feel confident in unit tests and have written some basic ones for the handlers. I did spend a little time wrestling with httptest because I haven't used it before.  

the top level routing behaviour is untested It could be nice to throw in some postman tests in a compose to run in CI

### documentation

becaues this is a go app I could utilize the godoc to extract documentation

### assumptions

- configuration is provided in ENV, eg: if being deployed on k8s the port should be available in the env. A 'sensible' default is included if this is absent.

- while we don't want 'description', 'version' and 'sha' hard-coded, I don't want them to be configurable at run-time either. They are "baked in" on build.
