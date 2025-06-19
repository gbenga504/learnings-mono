import React, { ReactNode, useContext } from "react";
import { createApiClient } from "./api";

const ApiContext = React.createContext(createApiClient());

export const ApiContextProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  return (
    <ApiContext.Provider value={createApiClient()}>
      {children}
    </ApiContext.Provider>
  );
};

export function useApi() {
  return useContext(ApiContext);
}
