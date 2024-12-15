# About GoRPS
GoRPS is a lightweight and high-performance tool built in Go using the Fiber framework to test server performance by simulating high RPS (Requests Per Second). This tool generates randomized 404 requests from a predefined list of paths to evaluate a server's ability to handle concurrent traffic effectively.

## Features

- **Dynamic Domain Configuration**: Easily update the target domain using a `.env` file.
- **Randomized Requests**: Generate and send requests from a list of 1000 unique paths.
- **High Concurrency**: Utilizes Go's goroutines to simulate real-world load.
- **Performance Insights**: Evaluate server RPS performance and identify bottlenecks.
- **Extensible**: Easily modify for additional HTTP methods, headers, or payloads.

## How It Works

1. **Load Configuration**:
   - The domain is loaded dynamically from the `.env` file.

2. **Path Generation**:
   - A file (`urls.txt`) contains 1000 unique paths (e.g., `/news1.php`, `/api/endpoint`, etc.).
   - Paths are read and randomized for each request.

3. **Concurrency**:
   - Requests are sent concurrently using goroutines, simulating real-world high traffic scenarios.

4. **Performance Evaluation**:
   - Measures the time taken to complete all requests, helping determine the server's RPS handling capacity.

## Example `.env` File

```env
DOMAIN=https://biwan.id
```

## Example `urls.txt`

```txt
/news1.php
/blog/random1.php
/api/data?query=test
/page123.html
/products/item99
...
```