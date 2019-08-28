# go-private-solution 

When working with golang private repos to be pull into a docker/container build process

Here is my RnD around the issue.

Solutions for each host repo provider are in the named directory. 

# Status 

1. github [working]
2. bitbucket [Issues]
3. gitlabs [Not started]

# Usage. 

1. cd into directory. `cd github.com`
2. run make target `make docker-up` this will build and run the container. 

# Further

1. Each directly has a `Dockerfile` that pull in the private repo code to be build and run. 
2. For this to work to your own use you will need to setup your own private repos
3. These solutions also use your local ssh key exposed via an `SSH_KEY_PATH` environment variable
4. The ssh key used requires no pass phase as you will not be able to add this into the containers build process. 


# Known Isuses

1. bitbucket is having the following issue during the build process: 

```
Step 11/13 : RUN go get -u bitbucket.org/lukekhamilton/private_golang
 ---> Running in a7dd446a2191
package bitbucket.org/lukekhamilton/private_golang: https://api.bitbucket.org/2.0/repositories/lukekhamilton/private_golang?fields=scm: 403 Forbidden
ERROR: Service 'app' failed to build: The command '/bin/sh -c go get -u bitbucket.org/lukekhamilton/private_golang' returned a non-zero code: 1
make: *** [Makefile:11: docker-build] Error 1
```

For the life of me I can't seam to find a solution for this currently. 
