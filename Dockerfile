FROM golang:1.19-rc-alpine

WORKDIR /app
#ENV GOROOT=/home/andrii/Downloads/go #gosetup
#ENV GOPATH=/home/andrii/go #gosetup

COPY ./ ./
#RUN go mod download

RUN go build -o /persons-web-micro

EXPOSE 8082

CMD [ "/persons-web-micro" ]