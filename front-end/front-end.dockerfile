FROM alpine:latest
RUN mkdir /app
COPY webApp /app
WORKDIR /app
CMD ["./webApp"]
