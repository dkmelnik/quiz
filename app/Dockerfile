FROM golang:1.18

ARG docker_id
ARG docker_user

RUN useradd -G sudo,root -u ${docker_id} -d /home/${docker_user} ${docker_user}
RUN mkdir /home/${docker_user}
RUN chown ${docker_user}:${docker_user} /home/${docker_user}

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz -C $(go env GOPATH)/bin
RUN curl -L https://github.com/danvergara/seeder/releases/download/v0.3.0/seeder_0.3.0_linux_amd64.tar.gz | tar xvz -C $(go env GOPATH)/bin
# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN apt-get update
RUN apt-get -y install net-tools lsof


WORKDIR /usr/src/app