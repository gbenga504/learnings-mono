const express = require("express");

const app = express();

app.get("/schools", function (_req, res) {
  return res.status(200).json({
    business: [
      { name: "Futa", country: "Nigeria" },
      { name: "MIT", country: "USA" },
    ],
  });
});

app.listen(4000, function () {
  console.log("Listening on port 4000");
});
