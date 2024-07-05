# Mechta Project

Этот проект включает два основных скрипта:

1. `cmd/generate/main.go` - генерирует файл `data/data.json` с миллионом объектов.
2. `cmd/sum/main.go` - вычисляет сумму всех чисел в `data/data.json` с использованием заданного числа горутин.

## Установка

### Клонирование репозитория
1. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/Mustafa0831/mechta.git
2. Перейдите в каталог проекта:
   ```sh
   cd mechta
### Инициализация Go модуля   
3. Инициализируйте Go модуль:
   ```sh
   go mod init mechta
### Установка зависимостей   
4. Установите все необходимые зависимости:
   ```sh
   go mod tidy
## Использование

### Генерация данных

Для генерации файла `data/data.json` выполните:

   ```sh
   go run cmd/generate/main.go
  ```
### Вычисление суммы

Для вычисления суммы всех чисел в `data/data.json` используйте следующий скрипт. Можно задать количество горутин для параллельной обработки данных через аргумент `-goroutines`.

Пример запуска с 4 горутинами:

   ```sh
   go run cmd/sum/main.go -goroutines=4
```
## Примеры запуска

   ```sh
   go run cmd/generate/main.go
   # data.json file has been generated successfully

   go run cmd/sum/main.go -goroutines=2
   # Total sum: -885
   # Time taken: 2.118001ms

   go run cmd/sum/main.go -goroutines=10
   # Total sum: -885
   # Time taken: 1.926039ms

   go run cmd/sum/main.go -goroutines=100
   # Total sum: -885
   # Time taken: 1.637317ms

   go run cmd/sum/main.go -goroutines=500
   # Total sum: -885
   # Time taken:  2.005726ms

   go run cmd/sum/main.go -goroutines=4
   # Total sum: -885
   # Time taken: 1.137947ms
```
Обратите внимание, что увеличение числа горутин не всегда приводит к линейному увеличению производительности из-за накладных расходов на управление горутинами и доступ к общей памяти.
# Тестирование
## Структура тестов
   ### pkg/data/data_test.go: Тесты для функций работы с данными.
   ### pkg/data/sum_processor_test.go: Тесты для процессора суммы.

## Запуск тестов
Для запуска тестов выполните:
```sh
go test -v ./pkg/...
```

## Запуск бенчмарков
Для запуска бенчмарков выполните:
```sh
go test -bench=. ./pkg/data
```
```sh
go test -bench=. ./cmd/sum
```
## Команды для работы с Docker
1. Создание Docker-образа:
```sh
docker build -t myapp .
```

2. Запуск Docker-контейнера:
```sh
docker run --rm -it -p 6060:6060 myapp -profile
```
3. ссылка для доступа к pprof:
```sh
http://localhost:6060/debug/pprof/
```
## Лицензия
Этот проект лицензирован под лицензией MIT. См. файл LICENSE для получения подробной информации.

## Вклад
Если вы хотите внести вклад в проект, пожалуйста, создайте форк репозитория и отправьте пул-реквест. Мы приветствуем любые улучшения и исправления.