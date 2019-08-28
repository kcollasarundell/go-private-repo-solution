FROM golang:latest

ARG SSH_KEY

WORKDIR /go/src

RUN mkdir -p /root/.ssh
RUN echo "$SSH_KEY" > /root/.ssh/id_rsa
RUN chmod 0600 /root/.ssh/id_rsa

RUN eval `ssh-agent` && ssh-add /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# Get private repo/s
RUN go get -u github.com/lukekhamilton/private_golang

# Remove ssh key
RUN rm /root/.ssh/id_rsa && unset SSH_KEY

ENTRYPOINT ["/go/bin/private_golang"]
