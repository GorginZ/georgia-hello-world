# GEORGIA-HELLO-WORLD


![main](https://github.com/gorginz/georgia-hello-world/actions/workflows/github-actions-release.yaml/badge.svg?branch=main)

registry: https://ghcr.io/gorginz/georgia-hello-world


### **required**: 

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


```docker run --rm -p 80:8001 -e HTTP_PORT=8001 georgia-hello-world```

----------

## publish new release
- update the VERSION accordingly (major.minor.patch)
- merge PR into main
- tag main with release version:

```git tag vx.x.x```
- push tag:

```git push origin vx.x.x```

- create a release. Can do this in github UI.
- select latest tag -> pipeline will build and push image to ghcr


## build and publish an image (non CI release)

auth with ghcr registry. e.g with PAK token:
>NOTE: this will require a georgia-can-I-have-a-gh-token/deploy-keys-request ;) 

``` echo $GHCR_TOKEN | docker login ghcr.io -u <user-name> --password-stdin  ```

run build and publish script

```./bin/build_and_publish.sh <environment> <image-tag>```


>NOTE: if on macOS with newer architecture you may need to explicitely set platform:

``` --platform linux/amd64 ```

Testing

```./bin/test.sh``` or ```go test -v ./...```


## Limitations

### CI/Pipeline stuff

I don't love how I read the description in the bin/build_and_publish.sh
this is probably fine for a small project - but for something bigger and with more configuration it would be worth utilizing some templating tooling to read various env configurations across environments.  

Everything is very coupled with github actions. But I have tried to put as much of the "work" in scripts as I can, because I think it's valuable to be able to easily run them locally. 
I thought of including another registry (ECR) for redundancy but it's not a requirement and I don't want to clog things up with extra yaml and scripts. 

### the code
I am not a go programmer, so while there are issues I am aware of there are likely many I am not.

One thing I'm not enthusiastic about is I have included the [gorrila/mux](https://github.com/gorilla/mux) pkg because I wanted a quick way to enforce what methods are allowed at the 'top' routing level. I notice they are looking for a new maintainer and think it's generally best not to rely on external tools for simple things like this, and I'm not sure it gives me a lot, but this was faster for me because I'm not a go-wiz.

### testing
 I like to rely and feel confident in unit tests and have written some basic ones for the handlers. I did spend a little more time than I'd hoped wrestling with ```httptest.recorder``` and some unexpected 200s. In particular the /status tests are a bit janky because the json encoder appends a newline but my string literal in the test's newline is being escaped.  
 I used a struct for representing the metadata because the original map was not in the order I wanted - my tests alerted me to this, so they've paid themselves off somewhat. 

The top level routing behaviour is untested It could be nice to throw in some postman tests in a compose to run in CI

### risks

- haven't included any code scanning tools in the pipeline, could use something like sonarcube, similarly the same can be said regarding image security.
- no branch protection for main

- currently gh actions is using a PAK (*only* scoped to packages - still bad though). This is very, very, very not reccomended! I will migrate to deploy keys. I wasn't aware until I read this: [gh actions security](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#considering-cross-repository-access)

### assumptions

- configuration is provided in ENV, eg: if being deployed on k8s the port should be available in the env. A 'sensible' default is included if this is absent.

- while we don't want 'description', 'version' and 'sha' hard-coded, I don't want them to be configurable at run-time either. They are "baked in" on build.
