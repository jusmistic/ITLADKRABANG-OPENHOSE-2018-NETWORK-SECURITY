#Network-101 deploy
cd $(pwd)/network-101-connect-to-the-world &&\
docker build -t net-101-container . &&\
cd ..

#Network-201 deploy
cd $(pwd)/network-201-QuickMath &&\
docker build -t net-201-container . &&\
cd .. &&\

echo "\n\n============[ Network Deploy Complete ]============\n\n"
