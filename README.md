# Мессенджер с системой оповещений

## Цель

Разработать систему мессенджера с поддержкой реального времени, интеграцией с системой оповещений и архитектурой микросервисов.

## Технологии

- Golang
- Gin
- Websocket
- gRPC
- PostgreSQL
- Redis
- Apache Kafka

## Основные требования

1. Пользователи могут регистрироваться, авторизовываться, восстанавливать пароль.
2. Пользователи могут отправлять и получать сообщения в реальном времени.
3. Поддержка групповых чатов.
4. Система оповещений о новых сообщениях, включая web-push уведомления.

## Архитектурные решения

- Используйте Gin для создания основного веб-API.
- Websocket для обеспечения мгновенной доставки сообщений.
- Используйте gRPC для внутренних вызовов между микросервисами.
- PostgreSQL для хранения основной информации пользователя и истории сообщений.
- Redis для кэширования часто запрашиваемой информации и управления сессиями.
- Apache Kafka для обработки и распределения оповещений.

## Микросервисы

- **Сервис авторизации**: Управляет процессами регистрации, авторизации и восстановления пароля. Использует PostgreSQL для хранения данных и Redis для управления сессиями.
- **Сервис сообщений**: Обрабатывает отправку, получение и хранение сообщений. Использует PostgreSQL для хранения и Websocket для мгновенной доставки.
- **Сервис оповещений**: Получает события о новых сообщениях из Kafka и отправляет уведомления пользователям через разные каналы (например, web-push, email).

## Процесс

1. При отправке сообщения, оно сохраняется в базе данных и через Websocket отправляется получателю.
2. Событие о новом сообщении также отправляется в Kafka.
3. Сервис оповещений слушает события из Kafka и отправляет уведомления соответствующим пользователям.

## Безопасность

1. Все данные пользователей, включая пароли, должны быть надежно защищены.
2. Обеспечьте шифрование данных при передаче между клиентом и сервером и между микросервисами.
3. Регулярно проводите аудит кода и инфраструктуры на предмет уязвимостей.

## Тестирование

1. Обеспечьте наличие юнит-тестов для основных компонентов системы.
2. Проведите нагрузочное тестирование для определения возможных узких мест и оптимизации производительности.