
# Tor IP Changer

## Overview
This project allows you to change your IP address using the Tor network. It automates the process of changing the IP using the `SIGNAL NEWNYM` command and verifies the new IP using the Tor SOCKS5 proxy.

## Requirements
- Go (Golang) installed
- Docker & Docker Compose installed
- Tor installed and running on your system (if not using Docker)
- curl installed for IP verification
- Basic understanding of terminal/command line usage

## Installation Instructions

### 1. Install Tor

#### On Windows:
1. Download and install Tor from the [Tor Project website](https://www.torproject.org/download/).
2. Follow the installation wizard and complete the installation.

#### On macOS:
Open a terminal and run the following commands:

```bash
brew install tor
```

#### On Linux (Ubuntu/Debian):
Open a terminal and run the following commands:

```bash
sudo apt update
sudo apt install tor
```

#### On Linux (CentOS/Fedora):
```bash
sudo dnf install tor
```

### 2. Configure Tor to Choose Exit Nodes

After installing Tor, you can configure it to select specific countries for your IP address.

#### On macOS/Linux:
1. Open a terminal and run the following command to edit the `torrc` file:

```bash
sudo nano /usr/local/etc/tor/torrc  # macOS
sudo nano /etc/tor/torrc            # Linux
```

#### On Windows:
1. Open `torrc` located in the Tor Browser's directory (e.g., `C:\Users\<your-username>\AppData\Roaming\tor\torrc`).

2. Add the following lines to select specific countries for your exit nodes:

```plaintext
ExitNodes {us},{ca},{de}     # Replace with desired country codes (US, Canada, Germany in this case)
ExcludeNodes {cn},{ru}       # Exclude these countries (China, Russia in this case)
StrictNodes 1                # Force using these exit nodes only
```

### 3. Restart Tor

After editing the configuration file, restart the Tor service:

#### On macOS:
```bash
brew services restart tor
```

#### On Linux:
```bash
sudo systemctl restart tor
```

#### On Windows:
Restart Tor Browser or run the following command if Tor is running as a service:

```cmd
net stop tor
net start tor
```

## Project Structure

- `cmd/main.go`: Main entry point for the project.
- `internal/tor`: Handles Tor connection and IP change logic.
- `internal/scheduler`: Manages scheduling of IP changes.
- `configs`: Configuration handling for scheduling intervals and Tor connection settings.
- `proxy`: Creates HTTP clients using the SOCKS5 proxy provided by Tor.

## Docker Setup

1. Build the Docker container:
```bash
docker-compose build
```

2. Run the Docker container:
```bash
docker-compose up
```

3. The application will change the IP every specified interval and display the new IP address.

## Optimized Build

For an optimized production build using Docker, you can add the `--build-arg` flag to optimize the Go binary. This reduces the binary size and enhances performance:

```bash
docker build --build-arg GO_BUILD_FLAGS='-ldflags="-s -w"' -t tor-ip-changer .
```

This will apply the optimization flags `-s` and `-w` to strip debugging information and reduce binary size.

## Example Output:
```bash
Current IP: 185.225.101.39
Successfully changed Tor IP
Current IP: 193.15.244.244
Successfully changed Tor IP
```

## Customization

You can edit the `config.yaml` file to change the interval of IP changes or customize the Tor control port.

Example `config.yaml`:
```yaml
tor:
  control_port: "127.0.0.1:9051" 
scheduler:
  interval: 60s  # Change IP every 60 seconds
```

## Troubleshooting

### 1. Tor is not starting
- Ensure Tor is installed and properly configured.
- Check the logs using the following commands:
  - On macOS/Linux: `tail -f /var/log/tor/log`
  - On Windows: Check the Tor Browser logs in the installation folder.

### 2. IP is not changing
- Make sure the Tor service is running correctly.
- Ensure your exit node configuration is properly set up in `torrc`.

### 3. Permission denied errors
- Make sure you have proper permissions to edit `torrc` and restart the Tor service. Use `sudo` if necessary on macOS/Linux.

## License
This project is open-source under the MIT License.
