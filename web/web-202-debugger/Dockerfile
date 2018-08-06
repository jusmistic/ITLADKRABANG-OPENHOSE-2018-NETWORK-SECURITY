FROM richarvey/nginx-php-fpm:1.5.3

WORKDIR /var/www/html

ENV PUID 1337
ENV PGID 1337
ENV WEBROOT /var/www/html/public

COPY ./srs /var/www/html/
RUN composer install

RUN find . -type f -exec chmod 644 {} \;
RUN find . -type d -exec chmod 755 {} \;
RUN chown -R 1337:1337 .
