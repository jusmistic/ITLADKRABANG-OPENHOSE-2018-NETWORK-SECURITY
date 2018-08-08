#Network-101 deploy
cd pwdned-101-easy-pwn &&\
docker build -t misc-101-container . &&\
docker run -dit --name misc-101 -p 7101:7101 misc-101-container &&\
cd ..

echo "\n\n============[ Misc Deploy Complete ]============\n\n"
