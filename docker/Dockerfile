# Step 1: Choose the right kitchen
FROM golang:1.22-alpine

# Step 2: Set up a mixing bowl
WORKDIR /app

# Step 3: Put the ingredients in the bowl
COPY ./codes/main.go .

# Step 4: Bake the cake (Compile the Go code)
RUN go build -o app-binary main.go

# Step 5: The final instruction
CMD ["./app-binary"]
