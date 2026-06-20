import java.io.*;
import java.net.*;

public class Main {
    private static final int PORT = 8080;
    private static final int BACKLOG = 10;
    private static final String RESPONSE =
        "HTTP/1.1 200 OK\r\n" +
        "Content-Type: text/plain\r\n" +
        "Content-Length: 13\r\n" +
        "Connection: close\r\n" +
        "\r\n" +
        "Hello, World!";

    private static void handleConnection(Socket client) throws IOException {
        try (client) {
            InputStream in = client.getInputStream();
            byte[] buf = new byte[4096];
            int n = in.read(buf);
            if (n <= 0) return;

            System.out.println("--- request ---");
            System.out.println(new String(buf, 0, n));

            OutputStream out = client.getOutputStream();
            out.write(RESPONSE.getBytes());
        }
    }

    public static void main(String[] args) throws IOException {
        try (ServerSocket server = new ServerSocket(PORT, BACKLOG)) {
            System.out.println("Listening on port " + PORT);
            while (true) {
                Socket client = server.accept();
                System.out.println("Connection from " + client.getRemoteSocketAddress());
                handleConnection(client);
            }
        }
    }
}
