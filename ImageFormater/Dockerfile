# Start from a base image with a desktop environment
FROM archlinux/base

# Install dependencies for building and running the Go application
RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm go xorg-server xorg-xinit xorg-xclock xterm libx11 libxinerama libxft && \
    pacman -Scc --noconfirm && \
    rm -rf /var/cache/pacman/pkg/*

# Set the working directory to where the Go application's code will be placed
WORKDIR /app

# Copy the Go application's source code into the container
COPY . .

# Build the application within the container
RUN go mod tidy && \
    go build -o app

# Set the entrypoint to launch the Go application within the container
ENTRYPOINT ["./app"]
