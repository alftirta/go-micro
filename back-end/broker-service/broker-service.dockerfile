# # build a base go docker image
# FROM golang:1.18-alpine AS builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o brokerService /cmd/api
# RUN chmod +x brokerService

# # build a tiny docker image
# FROM alpine:latest
# RUN mkdir /app
# COPY --from=builder /app/brokerService /app
# WORKDIR /app
# CMD ["./brokerService"]

FROM alpine:latest
RUN mkdir /app
COPY brokerService /app
WORKDIR /app
CMD ["./brokerService"]
