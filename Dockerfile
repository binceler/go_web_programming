FROM golang:1.21
WORKDIR go/src/goweb
COPY . .
CMD ["go/src/goweb/go_web_programming"]