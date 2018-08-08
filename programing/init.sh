#Prog-201 deploy
cd prog-201-buy-me-more &&\
docker build -t prog-201-container . &&\
docker run -dit --name prog-201 -p 10201:9000 prog-201-container &&\
cd ..


echo "\n\n============[ Programming Deploy Complete ]============\n\n"
