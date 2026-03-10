FROM php:8.3-apache

USER root

RUN apt-get update
RUN apt-get install -y libzip-dev
RUN apt-get install -y unzip

# INSTALL PHP EXTENSION INSTALLER
ADD https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions /usr/local/bin/
RUN chmod +x /usr/local/bin/install-php-extensions

# INSTALL PHP EXTENSIONS
RUN install-php-extensions intl
RUN install-php-extensions redis
RUN install-php-extensions opcache
RUN install-php-extensions memcached
RUN install-php-extensions pdo
RUN install-php-extensions mysqli
RUN install-php-extensions pdo_mysql

RUN a2enmod rewrite
RUN a2enmod ssl

COPY apache/vhost.conf /etc/apache2/sites-enabled/000-default.conf
