FROM alpine:latest
RUN mkdir /app
COPY loggerService /app
WORKDIR /app
CMD ["./loggerService"]
