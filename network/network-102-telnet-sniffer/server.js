var net = require('net');

var server = net.createServer(socket => {
    let eiei = {user : "admin", password: "Flag{Wendy_is_$0_Cut3}"};
    let system  = {state: "user", user: "admin", password: "Flag{Wendy_is_$0_Cut3}" };

    socket.on('connection', ()=>{

    });

    system.state = "user";
    socket.write(" === [ Welcome to Irene's System ] ===\r\n");
    socket.write(" ============ Please Login ===========\r\n");
    socket.write("Username: ");

    socket.on('data', (data) => {
        if(system.state === "user"){
            system.user = data.toString('ascii').trim();
            system.state = 'password';
            console.log('user ' + data);
            socket.write("Password: ");
        } else if(system.state === "password"){
            system.password = data.toString('ascii').trim();
            console.log('password ' + system.password);
            if (system.user === eiei.user && system.password === eiei.password){
                socket.write("Correct");
                socket.destroy();
            } else{
                socket.write("Incorrect");
                socket.destroy();
            }
        }
    });
    


    socket.on('error', function (err) {
        console.error(err.stack);
      });
});

server.listen(9102, '0.0.0.0');