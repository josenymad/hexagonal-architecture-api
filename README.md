# Hexagonal Architecture API

I developed this to improve my understanding of hexagonal architecture.
It's an API written in Go which connects to MongoDB.

The advantages of using hexagonal architecture is that the core logic of your application is de-coupled from any external services, which makes it more flexible and scalable if a migration to another external service would take place. For example, swapping from MongoDB to DynamoDB.
I found it to be far more complex than MVC architecture, but in the long run it means the whole project does not have to be re-written in circumstances such as the example above.

[This](https://tsh.io/blog/hexagonal-architecture/) is a great article which explains it in more detail.

This project is an extension of a [smaller mock API](https://github.com/josenymad/one-two-one-two) I built, only this will write the JSON data it receives to MongoDB.

I have also added a second route, to read all of the records in the database.

There isn't really any real world application for this program, I just wanted to build something to help me understand hexagonal architecture.

Going forward, I'd like to improve the error handling.
