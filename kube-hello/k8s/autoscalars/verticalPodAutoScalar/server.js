const express = require("express");

const app = express();

const fibonacci = (number) => {
  if (number === 1) return 1;

  if (number === 0) return 0;

  return fibonacci(number - 1) + fibonacci(number - 2);
};

app.get("/cpu", function (_req, res) {
  const result = fibonacci(44);

  res.status(200).json({ status: true, result });
});

app.get("/memory", function (_req, res) {
  new Array(20000000).fill(0);

  res.status(200).json({ status: true });
});

app.listen(3000, async function () {
  console.log("Listenning on port 3000");
});
