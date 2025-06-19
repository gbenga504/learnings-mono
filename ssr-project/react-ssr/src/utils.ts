import { matchRoutes } from "react-router-dom";
import { routes } from "./app/routes";
import { ICreateApiClient } from "./api/api";

export const loadPageDatas = async ({
  path,
  api,
}: {
  path: string;
  api: ICreateApiClient;
}): Promise<Record<string, any>> => {
  const matchedRoutes = matchRoutes(routes, path);
  let pageDatas = {};

  if (matchedRoutes) {
    for (const matchedRoute of matchedRoutes) {
      if (matchedRoute.route.loadData) {
        const result = await matchedRoute.route.loadData({ api });

        pageDatas = { ...pageDatas, [matchedRoute.route.id]: result };
      }
    }
  }

  return pageDatas;
};
