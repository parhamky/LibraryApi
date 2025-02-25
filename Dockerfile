FROM alpinelinux/golang AS build

WORKDIR /app

COPY . .

ENV foo /app/.env

USER root
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN  go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /LibraryApi ./cmd/main.go


FROM alpinelinux/golang

WORKDIR /

COPY $foo .

COPY --from=build /LibraryApi /LibraryApi

ENTRYPOINT [ "/LibraryApi" ]