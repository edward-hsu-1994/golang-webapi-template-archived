ARG MODE=release
ARG PORT=8080

FROM alpine:3.17.0 AS base
ENV GIN_MODE=$MODE
ENV PORT=$PORT
EXPOSE $PORT
RUN mkdir -p /app
WORKDIR /app





FROM golang AS build
#ENV GOPRIVATE=
WORKDIR /src

COPY . .

RUN rm -f -r ./dist
RUN mkdir -p ./dist
RUN sh ./scripts/update-swagger.sh
RUN go mod download
RUN go build -o ./dist/app ./cmd 
RUN cp -r ./assets ./dist
RUN cp -r ./configs ./dist





FROM base AS final
WORKDIR /app
COPY --from=build /src/dist .
RUN chmod +x ./app
ENTRYPOINT ["./app"] 