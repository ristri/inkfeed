# Inkfeed

Inkfeed is a lightweight, alternative Reddit client specifically designed for e-ink devices like Kindle. While optimized for e-readers, it works seamlessly across mobile devices and desktop browsers, providing a distraction-free browsing experience.

## Key Features

- No JavaScript for core functionality
- E-ink friendly interface with high contrast and clear typography
- Responsive design that adapts to different screen sizes
- Self-hostable on various platforms
- Uses Reddit's JSON API for content fetching

## Screenshots
<div align="center">
  <img src="./screenshots/inkfeed_kindle_sub.jpg?raw=true" alt="/r/cricket on a Kindle 10th generation" style="max-width: 100%; height: auto;">
  <p style="text-align: center; font-style: italic; margin-top: 10px; margin-bottom: 30px;">
    /r/cricket on a Kindle 10th generation
  </p>
</div>
<div align="center">
  <img src="./screenshots/inkfeed_kindle_comments.jpg?raw=true" alt="comments section on a Kindle 10th generation" style="max-width: 100%; height: auto;">
  <p style="text-align: center; font-style: italic; margin-top: 10px; margin-bottom: 30px;">
    Comments section on a Kindle 10th generation
  </p>
</div>
<div align="center">
  <img src="./screenshots/inkfeed_android_einkbro.jpg?raw=true" alt="/r/technology on an Android on EinkBro browser" style="max-width: 100%; height: auto;">
  <p style="text-align: center; font-style: italic; margin-top: 10px; margin-bottom: 30px;">
    /r/technology on an Android on EinkBro browser
  </p>
</div>

## Installation and Setup

### Prerequisites

- Go 1.21 or later
- Git

### Standard Installation

1. Clone the repository:
```bash
git clone https://github.com/ristri/inkfeed.git
cd inkfeed
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Development with Air (Hot Reload)

For development, you can use Air for hot reloading. To install Air, visit: [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)

Once Air is installed, simply run:
```bash
air
```

### Docker Installation

1. Build the Docker image:
```bash
docker build -t inkfeed .
```

2. Run the container:
```bash
docker run -p 8080:8080 inkfeed
```

Access the application at `http://localhost:8080`

## Self-Hosting Guide

### Standard Server
1. Build the binary:
```bash
go build -o inkfeed
```

2. Run the server:
```bash
./inkfeed
```

### Raspberry Pi
1. Install Go on Raspberry Pi:
```bash
sudo apt update
sudo apt install golang
```

2. Follow the standard installation steps above.

3. To run as a service, create a systemd service file:
```bash
sudo nano /etc/systemd/system/inkfeed.service
```

Add the following content:
```ini
[Unit]
Description=Inkfeed Reddit Client
After=network.target

[Service]
ExecStart=/path/to/inkfeed
WorkingDirectory=/path/to/inkfeed/directory
User=pi
Restart=always

[Install]
WantedBy=multi-user.target
```

Enable and start the service:
```bash
sudo systemctl enable inkfeed
sudo systemctl start inkfeed
```

### Android (Termux)
1. Install Termux from F-Droid 

2. Install required packages:
```bash
pkg update
pkg install golang git
```

3. Follow the standard installation steps

4. To keep the server running in background:
```bash
tmux
./inkfeed
# Press Ctrl+b then d to detach
```

## Important Notes

- This project uses Reddit's JSON API which has rate limits. Due to these limitations, no public instance is hosted.
- It's recommended to self-host for personal use only.
- The application is designed to work without JavaScript, making it ideal for e-readers.
- Some features may require JavaScript, but core browsing functionality works without it.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the [GPLv3](./LICENSE).
