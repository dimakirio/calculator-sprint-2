# calculator-sprint-2
#Установка и запуск
Для запуска проекта выполните следующие шаги:

Склонируйте репозиторий:
git clone https://github.com/Andreyka-coder9192/calc_go.git
cd calc_go

Убедитесь, что Go установлен и находится в $PATH (проверить версию можно командой go version).

Запустите API-сервер:

go run ./cmd/main.go
Сервер запустится на порту 8080 по умолчанию. Если необходимо изменить порт, установите переменную окружения PORT перед запуском:

В Linux и macOS
export PORT=9090
go run ./cmd/main.go
В Windows (PowerShell)
$env:PORT=9090
go run ./cmd/main.go
Использование API
Эндпоинт
POST /api/v1/calculate
Заголовки
Content-Type: application/json
Тело запроса
Пример:

{
  "expression": "2+2*2"
}
Ответы
Успешный запрос

Статус-код: 200 OK
Пример ответа:

{
  "result": "6"
}
Ошибка обработки выражения

Статус-код: 422 Unprocessable Entity
Пример ответа:

{
  "error": "Error calculation"
}
Неподдерживаемый метод

Статус-код: 405 Method Not Allowed
Пример ответа:

{
  "error": "Wrong Method"
}
Некорректное тело запроса

Статус-код: 400 Bad Request
Пример ответа:

{
  "error": "Invalid Body"
}
