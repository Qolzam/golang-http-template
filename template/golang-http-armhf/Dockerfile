 # BUILD* and TARGET* args are filled automatically                                                                          
 # BUILD* means the platform of the node doing the build                                                                     
 # TARGET* is the target platform of the build job                                                                           
 # The values are in global scope and do not leak to your commands                                                           
                                                                                                                             
 # this stage is fixed to worker platform (linux/amd64 in this case)                                                         
 FROM --platform=$BUILDPLATFORM golang:alpine AS build                                                                       
 # expose the values to the stage                                                                                            
 ARG TARGETOS                                                                                                                
 ARG TARGETARCH                                                                                                              
 # configure go to use target configuration                                                                                  
 ENV GOOS=$TARGETOS GOARCH=$TARGETARCH                                                                                       
 WORKDIR /go/src/github.com/tonistiigi/example                                                                               
 COPY main.go .                                                                                                              
 RUN go build -o /bin/hello .                                                                                                
                                                                                                                             
 FROM scratch AS release                                                                                                     
 COPY --from=build /bin/hello . 
                                                                                        
FROM binaries AS release                                                                                         
ENTRYPOINT ["/hello"]  