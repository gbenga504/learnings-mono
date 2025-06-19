import React from "react";
import { hydrateRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import { createApiClient } from "../api/api";

(function setupClient() {
  const pageDatas = JSON.parse(
    document.getElementById("appData")?.textContent!
  );

  const api = createApiClient();

  hydrateRoot(
    document.getElementById("root")!,
    <BrowserRouter>
      <App pageDatas={pageDatas} api={api} />
    </BrowserRouter>
  );
})();
