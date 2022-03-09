# если приложение разворачивается в докере изменить строку подключения к базе данных на :
DATABASE_URL="host=localhost user=user password=123 dbname=postgres port=5432 sslmode=disable"

# Описание эндпоинтов

## Создание чата 

POST localhost:8080/create 

{
"chat_name": "SuperChat",
"chat_author": "Nobody"
}

Если автор чата не чилится в юзерах,создает и чат и юзера

## Добавить сообщение 

POST localhost:8080/add-message

{
"message_author": "Nobody",
"chat_name": "SuperChat",
"message_text" : "Это сообщение из чата SuperChat"
}

Если автор сообщения не чилится в юзерах,создает и сообщение и юзера

## Получить id сообщений с глубиной поиска 

GET localhost:8080/chats/SuperChat/4 

Ответ :

{
"data": {
"0": "510510d6-4568-4b27-a66c-1735b8d6f0ba",
"1": "8bec377a-d3b5-489e-a667-b08a733d1159",
"2": "8db2f5c8-7577-4db7-bc4c-bc1182f5260b"
},
"success": true
}

## Получить информацию о сообщение по его id 

localhost:8080/chat/510510d6-4568-4b27-a66c-1735b8d6f0ba

Ответ :

{
"data": {
"ID": "510510d6-4568-4b27-a66c-1735b8d6f0ba",
"ChatID": "e83c624a-fd65-4baf-8595-cf20d610046a",
"message_author": "Nobody",
"chat_name": "SuperChat",
"message_text": "Это сообщение из чата SuperChat",
"CreatedAt": "2022-03-09T12:05:42Z"
},
"success": true
}