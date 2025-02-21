## Overview

This repository is of billing service

## Endpoints

Method | Path                               | Description                                   |                                                                         
---    |------------------------------------|------------------------------------------------
GET    | `/health`                          | Health page                                   |
GET    | `/metrics`                         | Страница с метриками                          |
GET    | `/v1/account/list`                 | Получение списка счетов пользователей         |
GET    | `/v1/account/get-balance/{userId}` | Получение баланса счёта по id пользователя    |
POST   | `/v1/account/create`               | Создание счёта пользователя                   |
PUT    | `/v1/account/buy`                  | Покупка (списание со счёта пользователя)      |
PUT    | `/v1/account/deposit`              | Депозит (пополнение счёта пользователя)       |