import type { AxiosInstance } from "axios";

export const auth = ({
  restClientForBackendServer,
  restClientForFrontendServer,
}: {
  restClientForBackendServer: AxiosInstance;
  restClientForFrontendServer: AxiosInstance;
}) => {
  return {
    login: ({ email, password }: { email: string; password: string }) => {
      return restClientForFrontendServer.post("/auth/login", {
        email,
        password,
      });
    },

    authenticate: ({
      email,
      password,
    }: {
      email: string;
      password: string;
    }) => {
      return restClientForBackendServer.post("/login", {
        email,
        password,
      });
    },
  };
};
