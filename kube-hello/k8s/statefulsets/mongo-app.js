const express = require("express");
const { MongoClient } = require("mongodb");

const app = express();

async function startDb() {
  const mongoClient = new MongoClient(
    "mongodb://admin:password@localhost:7000/test?replicaSet=rs0"
  );

  let client;

  try {
    client = await mongoClient.connect({ useUnifiedTopology: true });
    const db = client.db();

    console.log("Successfully connected to the client and database");

    return { db, client };
  } catch (err) {
    console.log("Something went wrong ===>", err);
  } finally {
    console.log("closing connection");
    client.close();
  }
}

async function main() {
  const { db } = await startDb();

  app.get("/todos", async function (_req, res) {
    const todos = await db.collection("todos").find();

    return res.status(200).json({ status: true, data: todos });
  });

  app.listen(9000, function () {
    console.log("Listening on port 9000");
  });
}

main();
