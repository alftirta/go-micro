FROM alpine:latest
RUN mkdir /app
COPY mailService /app
COPY templates /app/templates
WORKDIR /app
CMD ["./mailService"]
