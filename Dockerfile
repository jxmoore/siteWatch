
FROM golang:1.13.4-stretch as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build .
ENTRYPOINT [ "./AvailTest" ]
CMD [ "-k", "a"]