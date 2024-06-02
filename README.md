# Описание проекта

## Автор
Егоров Артём
telegram: @artemegorov06

## Запуск решения (Сборка исходников)

1. `git clone git@github.com:ArtemSoftware2006/Yandex.FinalProject.git` - клогирование репозитория
2. `cd /calculator` - переходим в рабочий каталог
3. `go run ./cmd/orchestrator/main.go` - запуск сервера с окрестраторосм
4. `go run ./cmd/agent/main.go` - запуск агентов

## Запуск решения (Docker)

### Версии

1. go version go1.21.3 linux/amd64

## Примеры использвоания

1. Добавление вычисления арифметического выражения
    ```shell
    curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
        "expression": "10 * 2"
        }'
    ```
    Ожидаемый ответ `{"id":"1717359192513988229"}`

2. Получение списка выражений
    ```shell
        curl --location 'localhost:8080/api/v1/expressions'
    ```

    Ожидаемый ответ: `{"expressions":[{"id":"1717359192513988229","status":"completed","result":20}]}`
 
 3. Получение выражения по его идентификатору:
    
    ```shell
        curl --location 'localhost:8080/api/v1/expressions/1717359192513988229'
    ```

    Ожидаемый ответ: `{"expression":{"id":"1717359192513988229","status":"completed","result":20}}`


## Переменные среды
Время выполнения операций задается переменными среды в милисекундах
 
**TIME_ADDITION_MS** - время выполнения операции сложения в милисекундах

**TIME_SUBTRACTION_MS** - время выполнения операции вычитания в милисекундах

**TIME_MULTIPLICATIONS_MS** - время выполнения операции умножения в милисекундах

**TIME_DIVISIONS_MS** - время выполнения операции деления в милисекундах
