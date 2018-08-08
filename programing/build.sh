#Prog-201 deploy
cd $(pwd)/prog-201-buy-me-more &&\
docker build -t prog-201-container . &&\
cd ..


echo "\n\n============[ Programming Deploy Complete ]============\n\n"
