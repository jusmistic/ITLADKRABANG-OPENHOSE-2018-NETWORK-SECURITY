#web-100 install 
cd $(pwd)/web-100-expore-the-web &&\
docker build -t web-100-container . &&\
docker run -dit --rm --name web-100 -p 8100:80 web-100-container &&
cd .. &&

#web-101 install 
cd $(pwd)/web-101-robot &&\
docker build -t web-101-container . &&\
docker run -dit --rm --name web-101 -p 8101:80 web-101-container &&
cd .. &&

#web-102 install
cd $(pwd)/web-102-Eltimate-robot &&\
docker build -t web-102-container . &&\
docker run -dit --rm --name web-102 -p 8102:80 web-102-container &&
cd .. &&

#web-201 install 
cd $(pwd)/web-201-CookieJar &&\
docker build -t web-201-container . &&\
docker run -dit --rm --name web-201 -p 8201:8201 web-201-container &&
cd .. &&

#web-202 install 
cd $(pwd)/web-202-debugger &&\
docker build -t web-202-container . &&\
docker run -dit --rm --env-file .env --name web-202 -p 8202:80 web-202-container 

echo "\n\n============[ Web Deploy Complete ]============\n\n"
