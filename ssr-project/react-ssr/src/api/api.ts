import axios from "axios";
import { auth } from "./auth";
import { dashboard } from "./dashboard";

export const createApiClient = ({
  authToken,
  rawHeaderCookies,
}: {
  authToken?: string;
  rawHeaderCookies?: string;
} = {}) => {
  const isDom = typeof window !== "undefined";
  const headersForBackendServer: { authToken?: string; cookie?: string } = {
    authToken,
  };

  if (!isDom && rawHeaderCookies) {
    headersForBackendServer.cookie = rawHeaderCookies;
  }

  const restClientForBackendServer = axios.create({
    baseURL: isDom ? "/api" : "http://localhost:6900",
    headers: headersForBackendServer,
  });

  // baseURL can be http://localhost:5555. '/' would work on client but not on server because server would default it to :80 (PORT) i.e default http TCP port
  const restClientForFrontendServer = axios.create({
    baseURL: isDom ? "/" : "http://localhost:5555",
  });

  return {
    auth: auth({ restClientForBackendServer, restClientForFrontendServer }),
    dashboard: dashboard({
      restClientForBackendServer,
      restClientForFrontendServer,
    }),
  };
};

export type ICreateApiClient = ReturnType<typeof createApiClient>;
