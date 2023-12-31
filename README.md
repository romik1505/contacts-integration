# unisender

## Предаврительная установка
Для запуска потребуется установка
- [ ] [docker](https://www.docker.com/products/docker-desktop/)
- [ ] [docker-compose](https://docs.docker.com/compose/install/)
- [ ] make - `sudo apt install build-essential`

## Перед запуском
- [ ] Изучите файл с переменными среды **_.env.example_**
- [ ] Создайте приватную интеграцию и указать uri интеграции в формате `https://URI/api/oauth/sign_in`
- [ ] Заархивирутей виджет из папки `./widget` указав в `script.js` url отправки токена
- [ ] Создайте файл **_.env_** и заполните данными полученными при создании интеграции.
  - [ ] Секретный ключ - API_SECRET
  
## Переменные среды .env
  - [ ] Параметры в верхней части файла (без префикса TEST_)
    - обрабатываются сервером по умолчанию в докере
- [ ] Параметры в нижней части файла (с префиксом TEST_) 
    - обрабатываются сервером если установлено _`ENVIRONMENT=local`_ для локального запуска
 
## Запуск в docker
_Перед запуском в докере убедитесь что установленные порты свободны для использования._
- [ ] В docker контейнере `make docker`

## Локальный запуск

- [ ] Локальный запуск `make run`

## Endpoints

Актуальное описание параметров и методов можно получить после запуска по ссылке [swagger](http://localhost:8080/swagger/index.html/) 


|                     Эндпоинт                     |             Описание              | 
|:------------------------------------------------:|:---------------------------------:|
|             `GET /api/account/{id}`              |      Информация об аккаунте       |
|         `GET /api/account/{id}/contacts`         |    Список контактов из amoCRM     |
| `GET /api/account/{id}/contacts/hook/unsub`      |      Отписка учетной записи       |
|       `GET /api/account/{id}/integrations`       |    Список интеграций аккаунта     |
|               `GET /api/accounts`                |    Список учетных записей         |
|             `GET /api/oauth/sign_in`             |        Добавление виджета         |
|            `POST /api/contacts/sync`             | Первичная синхронизация контактов |
|      `POST /api/account/{id}/contacts/hook`      |  Хук измений контактов аккаунта   |

