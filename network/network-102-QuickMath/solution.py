"""
    Solution for net 102
"""
import socket

def main():
    """ Socket Client to solve net 102 """

    #Establish Connection
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    #Assign Value
    host = '127.0.0.1'
    port = 9102

    # Connect to remote server
    sock.connect((host, port))                 

    # loop recv data
    while 1:
        #decode byte -> str
        msg = sock.recv(1024).decode('ascii')
        # print(msg)
        if 'or' not in msg and 'ar' not in msg and 'la' not in msg and 'ime' not in msg: #check trash text
            temp = msg.rstrip(" = ?\r\n").split(" + ")
            val = str(int(temp[0])+int(temp[1]))
            sock.send(val.encode())
        if 'la' in msg:
            print(msg)
            sock.close()
            return 'eiei'
        if 'ime' in msg:
            print("Timeout ---> Reconnecting")
            sock.close()
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.connect((host, port))                 


    sock.close()

main()
