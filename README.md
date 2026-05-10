# Go Docker Starter Project

A lightweight, containerized Go web application designed for self learning purpose. This project uses Docker Compose and volume mounts to create a seamless, "PHP-like" development experience.
## 📂 Project Structure

```text
my-project/
├── .gitignore              # Files and folders ignored by Git
├── README.md               # This documentation file
├── docker/
│   ├── .env                # Environment variables (build context, ports)
│   ├── docker-compose.yml  # Docker Compose orchestration
│   └── Dockerfile          # Instructions to build the Go environment
└── codes/
    └── main.go             # The actual Go application code
```

## Local Development
1. **Clone the Repository**: Start by cloning this repository to your local machine.
   ```bash
   git clone https://github.com/shakibmostahid/go-playground.git
   cd go-playground
    ```
2. Copy the env example files.
   ```bash
   cp docker/.env.example docker/.env
   ```
3. First time build with volume mount:
   ```bash
   cd docker
   docker compose build app
   docker compose run --rm app sh -c "go build -o app-binary main.go"
   ```
4. **Start the Docker Container**: Use Docker Compose to run the container.
    ```bash
    docker compose up -d app
    ```
5. Test the application by visiting `http://localhost:8080` in your web browser. You should see a welcome message confirming that the Go application is running.

