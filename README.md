# Hexagonal Architecture API

I developed this to improve my understanding of hexagonal architecture.
It's an API written in Go which connects to MongoDB.

The advantages of using hexagonal architecture is that the core logic of your application is de-coupled from any external services, which makes it more flexible and scalable if a migration to another external service would take place. For example, swapping from MongoDB to DynamoDB.
It does make it more complex, but in the long run it means the whole project does not have to be re-written in circumstances such as those.

[This](https://tsh.io/blog/hexagonal-architecture/) is a great article which explains it in more detail.
