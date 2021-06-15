FROM alpine as build

RUN apk add go git
RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN ls -l
RUN go build -o /main bootstrap.go

FROM alpine
COPY --from=build /main /main

ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie
COPY entry.sh /
RUN chmod 755 /entry.sh

CMD [ "/entry.sh", "/main" ]


