const { ApolloServer }  = require('apollo-server')
const { ApolloGateway } = require('@apollo/gateway')

const gateway = new ApolloGateway({
    serviceList: [
      { name: 'sales', url: 'http://localhost:8001/query' },
      { name: 'order', url: 'http://localhost:8002/query' },
    ],
  });
  
  const server = new ApolloServer({
    gateway,
    subscriptions: false,
  });
  
  server.listen().then(({ url }) => {
    console.log(`Server ready at ${url}`);
  });