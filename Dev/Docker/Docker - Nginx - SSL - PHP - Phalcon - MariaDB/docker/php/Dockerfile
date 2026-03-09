FROM phalconphp/cphalcon:v5.9.2-php8.4

USER root

RUN set -eux && install-php-extensions intl
RUN set -eux && install-php-extensions pdo
RUN set -eux && install-php-extensions pdo_mysql
RUN set -eux && install-php-extensions memcached
RUN set -eux && install-php-extensions opcache
RUN set -eux && install-php-extensions redis

WORKDIR /var/www/html

CMD ["php-fpm"]
