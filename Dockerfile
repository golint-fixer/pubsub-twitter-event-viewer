FROM golang:1.9.1 as builder

WORKDIR /go/src/github.com/mchmarny/simple-chat/
COPY . .

# restore to pinnned versions of dependancies 
RUN go get -u github.com/tools/godep
RUN godep restore

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o simple-chat \
    -tags netgo -installsuffix netgo .

# build the clean image
FROM scratch as runner
# copy the app
COPY --from=builder /go/src/github.com/mchmarny/simple-chat/simple-chat .
# copy static artifacts 
COPY --from=builder /go/src/github.com/mchmarny/simple-chat/static /static
COPY --from=builder /go/src/github.com/mchmarny/simple-chat/templates /templates

ENTRYPOINT ["/simple-chat"]
