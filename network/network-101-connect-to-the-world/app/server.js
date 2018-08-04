var net = require('net');
var server = net.createServer(socket => {
    console.log("Connection establish....\n");
    socket.on('end', () => {
        console.log('disconnected...');
    });
    socket.write("Flag{Hello_the_network_sector}\r\n");
});

server.listen(9101, '0.0.0.0');