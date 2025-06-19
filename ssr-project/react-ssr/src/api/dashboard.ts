import type { AxiosInstance } from "axios";

export const dashboard = ({
  restClientForBackendServer,
  restClientForFrontendServer,
}: {
  restClientForBackendServer: AxiosInstance;
  restClientForFrontendServer: AxiosInstance;
}) => {
  return {
    getAllData: async () => {
      const result = await restClientForBackendServer.get("/dashboard");

      return result.data;
    },
  };
};
