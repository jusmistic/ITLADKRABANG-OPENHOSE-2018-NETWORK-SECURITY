docker run -dit --rm --name web-100 -p 8100:80 web-100-container &
docker run -dit --rm --name web-101 -p 8101:80 web-101-container &
docker run -dit --rm --name web-102 -p 8102:80 web-102-container &
docker run -dit --rm --name web-201 -p 8201:8201 web-201-container &
docker run -dit --rm --env-file web-202-debugger/.env --name web-202 -p 8202:80 web-202-container &
docker run -dit --rm --name web-301 -p 8301:80 end1an/webvul