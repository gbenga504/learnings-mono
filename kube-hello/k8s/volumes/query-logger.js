const path = require("node:path");
const express = require("express");
const os = require("node:os");
const fs = require("node:fs");

const app = express();
const logDirectoryPath = path.join(__dirname, "./files");
const logFilePath = path.join(logDirectoryPath, "./log.json");

// function createLogFileIfNotExist() {
//   if (!fs.existsSync(logDirectoryPath)) {
//     fs.mkdirSync(logDirectoryPath);

//     console.log({
//       msg: "Created the directory",
//     });
//   }

//   if (!fs.existsSync(logFilePath)) {
//     fs.writeFileSync(logFilePath);

//     console.log({
//       msg: "Created the file",
//     });
//   }
// }

app.get("/", function (req, res) {
  try {
    console.log({
      msg: "Received request",
      payload: JSON.stringify(req.query),
    });
    const query = req.query;

    // createLogFileIfNotExist();

    const content = JSON.parse(fs.readFileSync(logFilePath, "utf-8") || "[]");

    console.log({
      msg: "Current content",
      payload: JSON.stringify(content),
    });

    const newEntry = { timeStamp: new Date().toISOString(), query };
    const newContent = [...content, newEntry];

    fs.writeFileSync(logFilePath, JSON.stringify(newContent), "utf-8");

    console.log({
      msg: "Processed request",
      payload: JSON.stringify(req.query),
    });

    return res.status(200).json({
      success: true,
      payload: newEntry,
    });
  } catch (error) {
    console.log("A criptic error occurred", error);

    res.status(400).json({
      success: false,
      payload: {
        errorName: error.name,
        errorMessage: error.message,
      },
    });
  }
});

app.get("/last-content", function (_req, res) {
  const content = JSON.parse(fs.readFileSync(logFilePath, "utf-8"));

  return res.status(200).json({
    success: true,
    payload: content.at(-1),
  });
});

app.get("/all-content", function (_req, res) {
  const content = JSON.parse(fs.readFileSync(logFilePath, "utf-8"));

  return res.status(200).json({
    success: true,
    payload: content,
  });
});

app.get("/crash", function () {
  setTimeout(function () {
    throw new Error("Something went wrong");
  }, 5);
});

app.listen(8000, async function () {
  console.log("Listenning on port 8000");
});
