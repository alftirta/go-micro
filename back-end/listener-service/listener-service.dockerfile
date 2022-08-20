FROM alpine:latest
RUN mkdir /app
COPY listenerService /app
WORKDIR /app
CMD ["./listenerService"]
