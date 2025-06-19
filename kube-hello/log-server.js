const path = require("node:path");
const express = require("express");
const os = require("node:os");
const fs = require("node:fs");

const app = express();
const logFilePath = path.join(__dirname, "./log.txt");

function sleep(duration = 1500) {
  return new Promise(function (resolve) {
    setTimeout(resolve, duration);
  });
}

function generateLogHash() {
  return Date.now().toString().split("").reverse().join("").substring(0, 5);
}

async function logFreeMemory() {
  const percentageOfFreeMemory = ((os.freemem() / os.totalmem()) * 100).toFixed(
    2
  );

  const message = `Log ${generateLogHash()}: Free memory has to be ${percentageOfFreeMemory}%\n`;
  fs.appendFileSync(logFilePath, message, "utf-8");

  await sleep(2000);
  logFreeMemory();
}

app.get("/", function (_req, res) {
  try {
    const fileContent = fs.readFileSync(logFilePath, "utf-8");
    const fileContentChunks = fileContent.trim().split(/\n/);
    const lastLog = fileContentChunks.at(-1) ?? "N/A";

    return res.status(200).json({
      success: true,
      payload: {
        lastLog,
      },
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

app.listen(8000, async function () {
  logFreeMemory();

  console.log("Listenning on port 8000");
});
