# BUILD* and TARGET* args are filled automatically                                                                          
# BUILD* means the platform of the node doing the build                                                                     
# TARGET* is the target platform of the build job                                                                           
# The values are in global scope and do not leak to your commands                                                              

FROM --platform=$BUILDPLATFORM openfaas/of-watchdog:0.7.3 as watchdog
FROM --platform=$BUILDPLATFORM golang:1.13-alpine3.11 as build
# expose the values to the stage                                                                                            
ARG TARGETOS                                                                                                                
ARG TARGETARCH                                                                                                              
# configure go to use target configuration                                                                                  
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH    

RUN apk --no-cache add git

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

ENV CGO_ENABLED=0

RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler
COPY . .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./function/vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

ARG GO111MODULE="off"
ARG GOPROXY=""

RUN go build --ldflags "-s -w" -a -installsuffix cgo -o handler .
RUN go test handler/function/... -cover

FROM --platform=$BUILDPLATFORM alpine:edge

# expose the values to the stage                                                                                            
ARG TARGETOS                                                                                                                
ARG TARGETARCH   

#RUN echo "I am using the default scratch"
# Add non root user and certs
RUN apk --no-cache add ca-certificates
RUN addgroup -S app
RUN adduser -S -g app app
RUN mkdir -p /home/app
RUN chown app /home/app

WORKDIR /home/app

COPY --from=build /go/src/handler/handler    .
COPY --from=build /usr/bin/fwatchdog         .
COPY --from=build /go/src/handler/function/  .

RUN chown -R app /home/app

USER app

ENV fprocess="./handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"

CMD ["./fwatchdog"]
