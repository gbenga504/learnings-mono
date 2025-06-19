const express = require("express");
const axios = require("axios");

const app = express();

// Create 2 applications talking to each other via kubernetes and load balance
// Also create apps using namespaces and talk to each other
app.get("/", async function (_req, res, next) {
  const API_URL = process.env.SCHOOL_API_FULL_URL;

  console.log("Request received", API_URL);

  try {
    const response = await axios.get(`${API_URL}/schools`);
    return res.status(200).json(response.data);
  } catch (error) {
    if (error instanceof axios.AxiosError) {
      console.log(
        "More on this ===>",
        error.stack,
        error.message,
        error.name,
        error
      );

      const message = error.response?.data || "Unknown error";

      return res.status(error.status || 500).json({ error: message });
    }

    next(error);
  }
});

app.use(function (error, _req, res, next) {
  if (res.headersSent) {
    return next(error);
  }

  console.log("Uncaught error", error);
  res.status(500).json({ error: "Error was not caught" });
});

app.listen(3000, function () {
  console.log("Listening on port 3000");
});
