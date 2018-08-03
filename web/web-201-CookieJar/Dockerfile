FROM node:8
WORKDIR /app
COPY app/ ./
COPY app/package*.json ./
RUN npm install
CMD npm start
EXPOSE 8201