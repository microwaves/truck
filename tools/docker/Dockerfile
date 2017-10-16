FROM golang:1.9 as builder

WORKDIR /go/src/github.com/microwaves/truck
ADD . .
RUN go get && CGO_ENABLED=0 go build -o /truck -a -tags netgo -ldflags '-w' main.go

FROM scratch
COPY --from=builder /truck /truck

ENTRYPOINT ["/truck"]
EXPOSE 4444
CMD ["--target", "$TRUCK_TARGET"]
