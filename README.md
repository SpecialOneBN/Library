# 📚 Library API

Проект REST API для управления библиотекой — книги, авторы и аренда книг.  
Реализован на языке Go с использованием принципов чистой архитектуры и PostgreSQL.

![CI](https://github.com/SpecialOneBN/library/actions/workflows/ci.yml/badge.svg)

---

## 🚀 Возможности

- 📖 Управление книгами (добавление, удаление, поиск)
- 🧑 Авторы и связи книг с авторами
- 🔄 Аренда и возврат книг
- 🧱 Чистая архитектура (слои: entity, repository, service/usecase, handler)
- 📜 Swagger-документация
- 🐘 PostgreSQL + Goose для миграций
- 🐳 Docker и `docker-compose`
- ⚙️ GitHub Actions CI (build → migrate → docker)

---

## 🧰 Технологии

- **Язык**: Go 1.24.1
- **БД**: PostgreSQL
- **Миграции**: Goose
- **Документация**: swaggo/swag
- **CI/CD**: GitHub Actions
- **Контейнеризация**: Docker

---

## 🛠️ Запуск локально

```bash
git clone https://github.com/SpecialOneBN/library.git
cd library
docker-compose up --build

Swagger будет доступен по адресу: http://localhost:8080/swagger/index.html
