FROM registry.wake-app.net/wakeapp/mysql:8.0

COPY files/initial.sql /docker-entrypoint-initdb.d/init.sql

USER mysql
