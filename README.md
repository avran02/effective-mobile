# effective mobile test task

## Запуск
```
git clone https://github.com/avran02/effective-mobile
cd effective-mobile
mv example.env .env
```
Затем заполнить .env файл настройками БД и поправить config.yml при необходимости
```
docker-compose up --build -d
```

## Использование

Спецификация доступна на localhost:8000/docs по умолчанию

Так же на localhost:8080 доступен adminer (UI к базе данных)

Миграции выполняются docker контейнером migrate/migrate

Сервис обогащающий пользовательские даннные лежит в папке enrich_user_data_service_mock и всегда возвращает замоканное значение