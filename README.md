#  Клиент для получения курсов криптовалют

### Для запуска:
```
go run ./cmd/app
```
### Использование:
Для получения курса определенной криптовалюты использовать запрос с ключом `name`:
<br>
**Request:**
```http
GET /currency?name=bitcoin
```
<br>

**Response:**
```
{
    "bitcoin": {
        "id":"bitcoin",
        "symbol":"btc",
        "name":"Bitcoin",
        "image":"https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1696501400",
        "current_price":62779,
        "market_cap":1235020018285,
        "last_updated":"2024-04-27T10:54:14.468Z"
    }
}
```

Для получения данных с внешнего API реализована гороутина, которая обновляет данные каждые 10 мин.