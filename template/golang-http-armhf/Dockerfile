 ### qolzam/qolzam-qc-test-go-fn:latest-master-1c589d5
 # kubectl run --image=qolzam/qolzam-qc-test-go-fn:latest-master-1c589d5 my-app --port=8080
 # BUILD* and TARGET* args are filled automatically                                                                          
 # BUILD* means the platform of the node doing the build                                                                     
 # TARGET* is the target platform of the build job                                                                           
 # The values are in global scope and do not leak to your commands                                                           
                                                                                                                             
 # this stage is fixed to worker platform (linux/amd64 in this case)                                                         
FROM --platform=linux/arm openfaas/of-watchdog:0.7.6 as watchdog
FROM --platform=$BUILDPLATFORM golang:alpine AS build                                                                       
# expose the values to the stage                                                                                            
ARG TARGETOS                                                                                                                
ARG TARGETARCH                                                                                                              
# configure go to use target configuration      

RUN apk --no-cache add git

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

ENV GOOS=$TARGETOS GOARCH=$TARGETARCH                                                                                       
WORKDIR /go/src/github.com/tonistiigi/example                                                                               
COPY main.go .                                                                                                              
RUN go build -o /bin/hello .                                                                                                
                                                                                                                            
FROM scratch AS binaries                                                                                                    
COPY --from=build /bin/hello .                                                                                              
COPY --from=build /usr/bin/fwatchdog .                                                                                              
                                                                                                                            
FROM binaries                                                                                        
ENTRYPOINT ["/hello"]   
