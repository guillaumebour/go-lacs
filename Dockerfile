FROM golang:1.19.0-buster AS build

COPY . src/lacs
WORKDIR src/lacs

RUN go mod download

RUN go build -o build/lacs cmd/lacs/main.go


FROM debian:stretch

RUN apt-get update && \
    apt-get install -y texlive-full && \
    apt-get install -y --no-install-recommends texlive-latex-recommended texlive-fonts-recommended && \
    apt-get install -y --no-install-recommends texlive-latex-extra texlive-fonts-extra texlive-lang-all && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /go/src/lacs/build/lacs /app/lacs
WORKDIR /app
EXPOSE 8081
ENV GIN_MODE=release

CMD ["./lacs"]
