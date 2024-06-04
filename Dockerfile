FROM golang 

WORKDIR app

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go get -d -v ./...

RUN make build

EXPOSE 8080

CMD ["./bin/app"]
