ARG database_url
FROM golang:1.21 AS builder

WORKDIR app

ENV DATABASE_URL=$database_url

RUN echo "This is the value for the arg ${database_url}"
RUN echo $DATABASE_URL

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/ cmd/
COPY database/ database/
COPY internal/ internal/
COPY public/ public/
COPY templates/ templates/
COPY Makefile .

RUN make build/prod

RUN make install
RUN make migrate/prod


FROM gcr.io/distroless/base-debian11 AS runner

WORKDIR /app
COPY --from=builder /app/dist ./
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ./web-app
