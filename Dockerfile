FROM golang:1.14

WORKDIR /Users/pankajsharma/go/src/totality-corp-assignment

COPY . .

RUN go get -v -t -d ./...
RUN go install -v ./...
EXPOSE 8001
CMD ["totality-corp-assignment"]