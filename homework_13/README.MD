Изучить документацию для API сервиса coincap.io (https://pro.coincap.io/api-docs)

Создайте аккуаунт, создайте API ключ и на ЯП Go сделайте GET запрос: /assets/. Вывод включает много полей. Выведите на экран ограниченным списком:
[{
    name: Bitcoin,
    priceUsd: 57435.2862859343831301,
},
{
    name: Ethereum,
    priceUsd: 3686.1846963999934439,
    
}
...
]

rest.conincap.io/v3/assets