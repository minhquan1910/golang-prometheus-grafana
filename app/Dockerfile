
FROM golang:1.21.4

# Installing OS dependencies here.
# Avoid using `apt` command outside this RUN block, will bloat image size.
RUN apt update -y && \
  apt install -y \
  make \
  postgresql-client && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

ENV ROOT /app

WORKDIR $ROOT

# Install air for live reload
RUN go install github.com/cosmtrek/air@v1.49

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8081

# Change to entrypoint after ECS config finish
CMD ["make", "run"]
