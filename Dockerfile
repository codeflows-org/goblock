FROM golang:1.14

COPY . .

RUN ./run_node.sh