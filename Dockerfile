FROM golang:1.15-alpine AS build

RUN mkdir -p /go/src/github.com/mohnishanandinc/bare-minimum-api/
#COPY main.go go.* /src/
ADD . /go/src/github.com/mohnishanandinc/bare-minimum-api/
WORKDIR /go/src/github.com/mohnishanandinc/bare-minimum-api/cmd
RUN ls -ltr
#COPY cmd/* /go/src/github.com/bare-minimum-api/
RUN CGO_ENABLED=0 go build -o /bin/bma

FROM scratch
COPY --from=build /bin/bma /bin/bma
ENTRYPOINT ["/bin/bma"]