const path = require("node:path");
const express = require("express");

const app = express();

app.get("/", function (_req, res) {
  return res.status(200).json({
    success: true,
    payload: {
      author: "Anifowoshe Gbenga David",
    },
  });
});

app.listen(8000, async function () {
  console.log("Listenning on port 8000");
});
