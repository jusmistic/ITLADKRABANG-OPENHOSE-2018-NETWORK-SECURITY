#Network-101 deploy
cd $(pwd)/pwned-101-easy-pwn &&\
docker build -t misc-101-container . &&\
cd .. && 

echo "\n\n============[ Misc Build Complete ]============\n\n"
