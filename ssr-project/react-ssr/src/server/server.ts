import express from "express";
import bodyParser from "body-parser";
import { createProxyMiddleware } from "http-proxy-middleware";
import cookieParser from "cookie-parser";

import { serverJsxEntry } from "../app/server-jsx-entry";
import { createApiClient } from "../api/api";

import authRoutes from "./auth-routes";

const app = express();

app.use(cookieParser());
app.use(bodyParser.json());

app.use(function (req, _res, next) {
  req.api = createApiClient({
    authToken: req.cookies.authToken,
    rawHeaderCookies: req.headers.cookie,
  });

  next();
});

app.use(
  "/api",
  createProxyMiddleware({
    target: "http://localhost:6900",
    pathRewrite: { "^/api": "/" },
    changeOrigin: true,
  })
);

app.use(authRoutes);

// Setting up route for static files(js,css,images...)
// Path should be relative to the main directory node runs from
app.use("/public", express.static("dist/public"));

app.get("/*", async function handleRequest(req, res) {
  if (req.url === "/favicon.ico") {
    return res.status(200).json({ status: "ok" });
  }

  const content = await serverJsxEntry(req);

  res.status(200).send(content);
});

app.listen(5555, () => {
  console.log("Listening on port 5555");
});
