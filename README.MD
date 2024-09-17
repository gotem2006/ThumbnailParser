# Thumbnail parser


## Вручную
Запуск:
- склонировать репозиторий в любую удобную директорию
```git clone https://github.com/sidecuter/sdo-python.git```
- перейти в директорию склонированного репозитория и установить утилиту
```go install client/main.go```
- запустить сервер
```go run cmd/app/main.go```
- воспользоваться утилитой
```client --async [urls]```
- example
```client --async https://www.youtube.com/watch?v=RMBfESIsOJo https://www.youtube.com/watch?v=RMBfESIsOJo```

## Запуск через docker-compose

Запуск:
- склонировать репозиторий в любую удобную директорию
```git clone https://github.com/sidecuter/sdo-python.git```
- перейти в директорию склонированного репозитория и установить утилиту
```go install client/main.go```
- запустить образ сервера
```docker-compose up```
- воспользоваться утилитой
```client --async [urls]```

## Запуск через makefile
- склонировать репозиторий в любую удобную директорию
```git clone https://github.com/sidecuter/sdo-python.git```
- воспользваться командой make run
