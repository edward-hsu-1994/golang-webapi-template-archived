ARG MODE=release
ARG PORT=8080

FROM alpine:3.17.0 AS base
ENV GIN_MODE=$MODE
ENV PORT=$PORT
EXPOSE $PORT
RUN mkdir -p /app
WORKDIR /app




FROM golang:alpine AS build
#ENV GOPRIVATE=
WORKDIR /app
RUN apk add --no-cache bash
#SHELL ["/bin/bash", "-c"]
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.8
RUN /bin/bash ./scripts/build.sh




FROM base AS final
WORKDIR /app
COPY --from=build /app/dist .
RUN ls -l
RUN chmod +x app
ENTRYPOINT ["./app"]
