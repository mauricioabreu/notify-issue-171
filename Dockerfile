FROM centos:6.8

RUN yum install -y git

# Installing go 1.9
RUN mkdir -p /usr/local/go \
  && curl -SLO "https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz" \
  && tar -xzf "go1.9.2.linux-amd64.tar.gz" -C /usr/local/go --strip-components=1 \
  && rm -f go1.9.2.linux-amd64.tar.gz

# Define go paths
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Build the binary with debug enable to read notify events
RUN go get -u github.com/rjeczalik/notify

COPY . $GOPATH
WORKDIR $GOPATH