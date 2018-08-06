var net = require('net');

var server = net.createServer(socket => {
    socket.setTimeout(5000);
    socket.write('Input Sometext here:');
    socket.on('data', data =>{
        if (data.length > 64){
            socket.write("Flag{Y0u__c@n__pwn__m3}\r\n");
            socket.destroy();
        }else{
            socket.write("Aww Can't Change?");
            socket.destroy();
        }
    });
    socket.on('error',()=>{
        console.error(err.stack);
    });
});

server.listen(7101, '0.0.0.0');