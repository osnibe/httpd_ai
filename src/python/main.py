import socket
import sys

PORT = 8080
BACKLOG = 10
RESPONSE = (
    b"HTTP/1.1 200 OK\r\n"
    b"Content-Type: text/plain\r\n"
    b"Content-Length: 13\r\n"
    b"Connection: close\r\n"
    b"\r\n"
    b"Hello, World!"
)


def handle_connection(conn, addr):
    with conn:
        data = conn.recv(4096)
        if not data:
            return
        print(f"--- request ---\n{data.decode(errors='replace')}")
        conn.sendall(RESPONSE)


def main():
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server:
        server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        server.bind(("", PORT))
        server.listen(BACKLOG)
        print(f"Listening on port {PORT}")

        while True:
            conn, addr = server.accept()
            print(f"Connection from {addr[0]}:{addr[1]}")
            handle_connection(conn, addr)


if __name__ == "__main__":
    main()
