FROM golang:1.16.3

RUN apt-get update && apt-get install -y jq

ADD https://github.com/go-swagger/go-swagger/releases/download/v0.26.1/swagger_linux_amd64 /tmp/swagger
RUN chmod +x /tmp/swagger && mv /tmp/swagger /usr/local/bin/swagger

ADD https://raw.githubusercontent.com/cosmtrek/air/master/install.sh /tmp/air.sh

RUN sh /tmp/air.sh -b /bin && rm -rf /tmp/air.sh

COPY .air.toml /root/.air.toml

WORKDIR /app

ENTRYPOINT [ "air", "-c", "/root/.air.toml"]
