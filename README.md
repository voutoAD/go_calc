# go_calc

# Проект по созданию для реализованного консольного </br> калькулятора серверной обертки

## 1.Условия составления выражения 
### Выражение должно состоять только из цифр(0-9), принимается только полная математическая нотация(без опуская знака * перед или после скобок и т.д.)

# 2. Использование
## Сервер собой представляет один endpoint API с адресом `/api/v1/calculate`. Обработка выражение происходит при запросе с методом `POST` и с телом запроса:
```json
{
  "expression": "<ваше выражение>"
}
```
## Формат ответа при корректных данных:
```json
{
  "result": "<число типа float с 6-ю знаками после запятой>"
}
```

## При не корректном вводе формат таков(статус 422):
```json
{
  "error": "Expression is not valid"
}
```
## При неверном методе запроса(статус 405):
```json
{
  "error": "Method not allowed"
}
```
## И при ошибке на сервере(статус 500):
```json
{
  "error": "Internal server error"
}
```

## 3. Запуск
### Запуск в режиме сервера со стандартными настройками:
```bash
go mod tidy
go build cmd/main.go
./main
```

### Для запуска в консольном режиме в [файле](cmd/main.go) уберите комментарий со строчки 10 и добавьте на 13 строчке комментарий. Команда для запуска та же самая.

## 4. Примеры работы:
### 4.1 Правильное тело запроса:
#### Запрос
```bash
curl --location 'localhost:3000/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
#### Ответ
```json
{
	"result": "6.000000"
}
```

### 4.2 Не правильное тело запроса:
#### Запрос
```bash
curl --location 'localhost:3000/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*-2"
}'
```
#### Ответ
```json
{
	"error": "Expression is not valid"
}
```

### 4.3 Пустое тело запроса:
#### Запрос
```bash
curl --location 'localhost:3000/api/v1/calculate' \
--header 'Content-Type: application/json' \
```
#### Ответ
```json
{
	"error": "Internal server error"
}
```
### 4.4 Не правильный метод запроса
#### Запрос
```bash
curl --location 'localhost:3000/api/v1/calculate'
```
#### Ответ
```json
{
	"error": "Method not allowed"
}
```