import React from "react";
import type { Request } from "express";
import { renderToString } from "react-dom/server";
import { StaticRouter } from "react-router-dom/server";
import App from "./App";
import serialize from "serialize-javascript";
import { loadPageDatas } from "../utils";

export const serverJsxEntry = async (req: Request) => {
  let pageDatas = await loadPageDatas({ path: req.path, api: req.api });

  const jsx = renderToString(
    <StaticRouter location={req.url}>
      <App pageDatas={pageDatas} api={req.api} />
    </StaticRouter>
  );

  return `
            <html>
                <head>
                    <title>Hello world</title>
                </head>
                <body>
                    <div id="root">${jsx}</div>
                    <script async src="/public/client.bundle.js"></script>
                    <script id="appData" type="application/json">${serialize(
                      pageDatas,
                      { isJSON: true }
                    )}</script>
                </body>
            </html>
        `;
};
