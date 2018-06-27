FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/mchmarny/pubsub-twitter-event-viewer/
COPY . .

# restore to pinnned versions of dependancies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o tview \
    -tags netgo -installsuffix netgo .

# build the clean image
FROM scratch as runner
# copy the app
COPY --from=builder /go/src/github.com/mchmarny/pubsub-twitter-event-viewer/tview .
# copy static artifacts
COPY --from=builder /go/src/github.com/mchmarny/pubsub-twitter-event-viewer/static /static
COPY --from=builder /go/src/github.com/mchmarny/pubsub-twitter-event-viewer/templates /templates

ENTRYPOINT ["/tview"]
