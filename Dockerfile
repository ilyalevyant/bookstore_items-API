FROM golang:1.17.0

ENV REPO_URL=github.com/ilyalevyant/bookstore_items-API
ENV GOPATH=/go
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src

#/Users/ilyalevyant/go/src/github.com/ilyalevyant/bookstore_items-API
COPY src $WORKPATH
WORKDIR $WORKPATH
RUN go build -o items-api .
EXPOSE 8081

CMD ["./items-api"]