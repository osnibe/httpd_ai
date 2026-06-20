'use strict';

const net = require('net');

const PORT = 8080;
const RESPONSE =
    "HTTP/1.1 200 OK\r\n" +
    "Content-Type: text/plain\r\n" +
    "Content-Length: 13\r\n" +
    "Connection: close\r\n" +
    "\r\n" +
    "Hello, World!";

const server = net.createServer((conn) => {
    const addr = `${conn.remoteAddress}:${conn.remotePort}`;
    console.log(`Connection from ${addr}`);

    conn.once('data', (data) => {
        console.log('--- request ---');
        console.log(data.toString());
        conn.end(RESPONSE);
    });
});

server.listen(PORT, () => {
    console.log(`Listening on port ${PORT}`);
});
