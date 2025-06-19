const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const cookieParser = require("cookie-parser");

const app = express();

app.use(cors());

app.use(bodyParser.json());
app.use(cookieParser());

const isAuthenticatedUser = (req, res, next) => {
  if (req.headers.authToken || req.cookies.authToken) {
    return next();
  }

  return res.status(401).json({ message: "UnAuthorized" });
};

app.post("/login", function (req, res) {
  const { email, password } = req.body;

  if (email === "johndoe@gmail.com" && password === "password") {
    return res.status(200).json({ authToken: "test" });
  }

  res.status(401).json({ message: "Incorrect login info" });
});

app.get("/dashboard", isAuthenticatedUser, function (req, res) {
  return res.status(200).json({
    data: [
      { id: 1, name: "Berlin" },
      { id: 2, name: "Hamburg" },
    ],
  });
});

app.get("/contract", function (req, res) {
  return res.status(200).json({
    name: "Wunderflats",
  });
});

app.listen(6900, () => {
  console.log("Listening on port 6900");
});
