# Homebroker Microservice

## To test the project

1. Run `docker-compose up -d`
2. Access `localhost:9021` to access Kafka Control Center
3. Create `orders` and `processed_orders` topics
4. Produce messages in `orders` topic

### Examples of orders

```json
{
    "order_id": "54321",
    "investor_id": "09876",
    "asset_id": "xyz",
    "current_shares": 200,
    "shares": 75,
    "price": 67.89,
    "order_type": "SELL"
}
```
```json
{
    "order_id": "67890",
    "investor_id": "12345",
    "asset_id": "xyz",
    "current_shares": 150,
    "shares": 75,
    "price": 67.89,
    "order_type": "BUY"
}
```
