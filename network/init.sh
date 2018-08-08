#Network-101 deploy
cd network-101-connect-to-the-world &&\
docker build -t net-101-container . &&\
docker run -dit --name net-101 -p 9101:9101 net-101-container &&\
cd ..

#Network-201 deploy
cd network-201-QuickMath &&\
docker build -t net-201-container . &&\
docker run -dit --name net-201 -p 9201:9201 net-201-container &&\
cd .. &&\

echo "\n\n============[ Network Deploy Complete ]============\n\n"
