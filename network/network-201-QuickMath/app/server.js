var net = require('net');

var server = net.createServer(socket => {
    socket.setTimeout(5000);
    let connection = {score: 0};

    function ask(){
        let num1 = Math.floor(Math.random()*10000);
        let num2 = Math.floor(Math.random()*10000);
        socket.write(num1+" + "+num2+" = ?\r\n");
        return num1+num2;

    };

    socket.write("[*]--- Start Game! ---[*]\r\n");
    let ans = ask();
    
    socket.on('data',(data) => {
        if (ans == parseInt(data)){
            connection.score += 1;
            console.log(connection.score+"\r\n");
            socket.write("[*]--- Correct!!! ---[*]\r\n");
            if (connection.score >= 10){
                socket.write("Flag{Qu1ckM@th_M4st3r}");
                socket.destroy();
            }
            ans = ask();
        } else{
            socket.write("[*]--- Incorrect!!! ---[*]\r\n");
            socket.destroy();
        }
    });

    socket.on('timeout', ()=>{
        socket.write("[*]--- Time OUT!!! ---[*]\r\n");
        socket.destroy();
    });
    socket.on('error', function (err) {
        console.error(err.stack);
      });
});

server.listen(9102, '0.0.0.0');