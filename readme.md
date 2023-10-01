# litongjava Go Simple HTTP Server

This is a simple HTTP server written in Go that serves static files from a specified directory.

## Features

- Serve static files from a directory.
- Specify the port to listen on.
- Display available IP addresses for accessing the server.

## Usage

### Downloading the Release

You can download the latest release of this project from the [GitHub Releases](https://github.com/litongjava/litongjava-go-simple-http-server/releases) page. Choose the release that matches your operating system and architecture.

### Running the Server

After downloading the release, extract the archive (if applicable) and change to the project directory. Then, run the HTTP server using the following command:

```shell
./http-server -port 3000 -dir /path/to/your/directory
```

Replace `/path/to/your/directory` with the path to the directory containing the files you want to serve. You can also specify a different port using the `-port` flag (default is 3000).

### Accessing the Server

After starting the server, you can access it in your web browser using one of the displayed IP addresses and the specified port. For example:

- http://127.0.0.1:3000 (localhost)
- http://192.168.0.100:3000 (local network)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
