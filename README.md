# stcTask

Пороли, переменные окружения и дам бд все тестовое или из мертвых проектов. И не является какой либо достоверной информацией.

Запуск контейнеров
```
make start
```

Запуск сервера отдельно
```
make run
```
Количество записей и запросов задается в docker-compose переменными окружения

Документация сервера находится в папке server/docks 
так же она прикреплена к роутеру и настроена для работы по адресу контейнера
```
http://localhost:8080/swagger/index.html
```

тесты не сделал так как они начали ломать логер да и вышла с ними заминка из за того что они падали даже несмотря на правильные полученые даные. К сожалению я сложнее юниттестов обычных функций на входные и выходные параметры не делал потому в проблеме разобраться не смог. 
