# music-playlist-api
В качестве протокола взаимодействия сервиса с клиентами используется gRPC

Модули
1. [Player](#player)
2. [Song](#song)
3. [Playlist](#playlist)

# Player
Модуль для обеспечения работы с плейлистом 

- Плейлист внутри реализован с использованием двусвязного списка
- Воспроизведение не блокирует методы управлнеия 
- Метод воспроизведения начинает воспроизведение с длительностью, ограниченной свойством Duration песни. Воспроизведение эмулируется длительной операцией.
- Следующая песня воспроизводится автоматически после окончания текущей песни.
- Метод Pause приостановливает текущее воспроизведение, и когда воспроизведение вызывается снова, оно продолжается с момента паузы.
- Метод AddSong добавит новую песню в конец списка.
- Вызов метода Next начинает воспроизведение следущей песни. Таким образом текущее спроизведение остановливается и начинается воспроизведение следущей песни
- Вызов метода Prev остановливает текущее воспроизведение и начинает воспроизведение предыдущей песни.
- Реализация метода AddSong проводиться с учетом одновременного, конкурентного доступа.
- Реализация методов Next/Prev проводится с учетом одновременного, конкурентного доступа.

Для хранения сущностей используется локальное хранилище (в проекте реализовано на мапе) 
## AddPlayer 
Создание проигрывателя 

##### Example pequest:
```json
{
    "playlistUuid": "b5af7ef2-7b10-4acb-9c15-6c9a946ec5d1"
}
```
##### Example response:
```json
{
    "playerUuid": "795e149a-651d-11ef-b4e2-00ff1b8d2930",
    "error": ""
}
```
## DeletePlayer 
Удаление проигрывателя

##### Example pequest:
```json
{
    "playerUuid": "52ae5c47-6486-11ef-962a-00ff1b8d293"
}
```
##### Example response:
```json
{
    "error": "player not exists"
}
```
##### Example response:
```json
{
    "error": ""
}
```

## AddSong 
Добавление песни в очередь воспроизведения (также обновится сущность связанного с плеером плейлиста)

##### Example pequest:
```json
{
    "playerUuid": "e7451576-64d3-11ef-b85a-00ff1b8d2930",
    "songUuid": "2e4fddf3-026f-44ca-9308-b4e1f5445bb5"
}
```
##### Example response:
```json
{
    "error": ""
}
```
## DeleteSong 
Удаление песни из очереди воспроизведения (также обновится сущность связанного с плеером плейлиста)

##### Example pequest:
```json
{
    "playerUuid": "e7451576-64d3-11ef-b85a-00ff1b8d2930",
    "songUuid": "2e4fddf3-026f-44ca-9308-b4e1f5445bb5"
}
```
##### Example response:
```json
{
    "error": ""
}
```
#Play 
Начать вопроизведение плейлиста
##### Example pequest:
```json
{
    "playerUuid": "36f056fa-651e-11ef-b4e2-00ff1b8d2930"
}
```

В качестве ответа клиенту приходит стрим
##### Example response:
```json
{"log":"playing nostrud: 15.000000/16.000000 seconds\n","error":""}
```

#Next 
Начать вопроизведение следующей песни

##### Example pequest:
```json
{
    "playerUuid": "36f056fa-651e-11ef-b4e2-00ff1b8d2930"
}
```
В качестве ответа клиенту приходит стрим
##### Example response:
```json
{"log":"playing nostrud: 15.000000/16.000000 seconds\n","error":""}
```

#Prev 
Начать вопроизведение предыдущей песни

##### Example pequest:
```json
{
    "playerUuid": "36f056fa-651e-11ef-b4e2-00ff1b8d2930"
}
```
В качестве ответа клиенту приходит стрим
##### Example response:
```json
{"log":"playing nostrud: 15.000000/16.000000 seconds\n","error":""}
```
#Pause 
Осстановить вопроизведение песни

##### Example pequest:
```json
{
    "playerUuid": "36f056fa-651e-11ef-b4e2-00ff1b8d2930"
}
```

##### Example response:
```json
{
    "error": ""
}
```




# Song
Модуль работы с песнями
Сущности хранятся в postgres 
## Create 
Создание песни 

##### Example pequest:
```json
{
  {
    "info": {
        "Duration": {
            "nanos": 98061907,
            "seconds": "16"
        },
        "title": "nostrud"
    }
}
```
##### Example response:
```json
{
    "uuid": "7d47865b-2922-46f6-a754-18f61f1086eb",
    "error": ""
}
```
## Get 
Получить песню

##### Example pequest:
```json
{
    "uuid": "2e4fddf3-026f-44ca-9308-b4e1f5445bb5"
}
```
##### Example response:
```json
{
    "song": {
        "uuid": "2e4fddf3-026f-44ca-9308-b4e1f5445bb5",
        "song_info": {
            "title": "elit do aliqua",
            "Duration": {
                "seconds": "9",
                "nanos": 0
            }
        }
    },
    "error": ""
}
```

## Delete 
Удаление песни 

##### Example pequest:
```json
{
    "uuid": "a927d481-cf75-4db6-8055-9fdfba4ed840"
}
```
##### Example response:
```json
{
    "error": "song not found sql: no rows in result set"
}
```
##### Example response:
```json
{
    "error": ""
}
```
## Update 
Обновление песни 

##### Example pequest:
```json
{
    "info": {
        "Duration": {
            "nanos": 1935276818,
            "seconds": "9"
        },
        "title": "elit do aliqua"
    },
    "uuid": "2e4fddf3-026f-44ca-9308-b4e1f5445bb5"
}
```
##### Example response:
```json
{
    "error": ""
}
```

# Playlist
Модуль работы с плейлистами
Сущности хранятся в postgres 
## Create 
Создание плейлиста 

##### Example pequest:
```json
{
    "info": {
        "name": "veniam consectetur quis cillum",
        "songList": [
            "2e4fddf3-026f-44ca-9308-b4e1f5445bb5",
            "7b88329d-7b68-4b9d-8138-d7e40f123997"
        ],
        "userId": "elit magna"
    }
}
```
##### Example response:
```json
{
    "uuid": "15e9d1ed-3cde-452e-bc9a-3f626d78920b",
    "error": ""
}
```
## Get 
Получие плейлиста 

##### Example pequest:
```json
{
    "uuid": "15e9d1ed-3cde-452e-bc9a-3f626d78920b"
}
```
##### Example response:
```json
{
    "playlist": {
        "uuid": "15e9d1ed-3cde-452e-bc9a-3f626d78920b",
        "playlist_info": {
            "songList": [
                "2e4fddf3-026f-44ca-9308-b4e1f5445bb5",
                "7b88329d-7b68-4b9d-8138-d7e40f123997"
            ],
            "name": "veniam consectetur quis cillum",
            "userId": "elit magna"
        }
    },
    "error": ""
}
```

## Delete 
Удаление плейлиста

##### Example pequest:
```json
{
    "uuid": "15e9d1ed-3cde-452e-bc9a-3f626d78920b"
}
```
##### Example response:
```json
{
    "error": ""
}
```
## Update 
Обновление плейлисат 

##### Example pequest:
```json
{
    "info": {
        "songList": [
                "51c85264-31d6-4480-b621-e67cb7d03a79",
                "6826397e-aeef-437e-92d7-15abfff01272"
            ],
        "name": "veniam consectetur quis cillum",
        "userId": "elit magnaqqqq"
    },
    "uuid": "15e9d1ed-3cde-452e-bc9a-3f626d78920b"
}
```
##### Example response:
```json
{
    "error": ""
}
```

