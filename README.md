# graphql-federation-go

Graphql Federation Implementation Example with Go

## Try it
### Clone project
```bash
git clone https://github.com/aeramu/graphql-federation-go.git
```
### Run sales-service
```bash
cd ./sales-service
env PORT=8001 go run .
```
### Run order-service
```bash
cd ./order-service
env PORT=8082 go run .
```
### Run gateway
```bash
cd ./gateway
yarn
yarn start
```
### Try on playground
open ```localhost:4000``` on your browser and try query from both services
```graphql
{
  getSales(id:"some-sales-id"){
    id
    name
    code
    phone
  }
  getOrder(id:"some-order-id"){
    id
    products
    totalAmount
    status
  }
}
```
