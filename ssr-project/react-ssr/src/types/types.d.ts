import type { RouteObject } from "react-router-dom";
import type { ICreateApiClient } from "../api/api";
import type { FC } from "react";

declare global {
  namespace Express {
    interface Request {
      api: ICreateApiClient;
    }
  }
}

declare module "react-router-dom" {
  type RouteObjectWithLoadData = RouteObject & {
    id: string;
    component: FC<any>;
    loadData?: ({ api: ICreateApiClient }) => Promise<any>;
  };
}
