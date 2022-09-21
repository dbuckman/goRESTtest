FROM registry.access.redhat.com/ubi8/go-toolset as build
COPY . .

# Build the Go app
RUN go mod download
# COPY *.go ./
RUN go build -o goRESTtest cmd/restTest/main.go

FROM registry.access.redhat.com/ubi8-micro
COPY --from=build /opt/app-root/src/goRESTtest .
CMD ./goRESTtest
