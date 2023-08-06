# Stage 1: build go
FROM golang as build

# Copy the source files from the host
COPY . /src

# Set the working dir
WORKDIR /src

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o product ./product

# Stage 2: build service
FROM scratch

# Copy the binaly from the build
COPY --from=build /src/product .

# Connect port to this Docker
EXPOSE 50051

# Run KVS
ENTRYPOINT ["/product"]