import React, { useEffect, useRef, useState } from "react";
import {
  matchRoutes,
  useLocation,
  renderMatches,
  RouteObjectWithLoadData,
} from "react-router-dom";

import { routes } from "./routes";
import { ApiContextProvider } from "../api/ApiContext";
import { loadPageDatas } from "../utils";
import { ICreateApiClient } from "../api/api";

interface IProps {
  pageDatas: Record<string, any>;
  api: ICreateApiClient;
}

const transformMatches = ({
  matches,
  pageDatas,
}: {
  matches: ReturnType<typeof matchRoutes<RouteObjectWithLoadData>>;
  pageDatas: Record<string, any>;
}): ReturnType<typeof matchRoutes<RouteObjectWithLoadData>> => {
  if (matches) {
    return matches.map((match) => {
      const Component = match.route.component;

      return {
        ...match,
        route: {
          ...match.route,
          element: <Component pageData={pageDatas[match.route.id]} />,
        },
      };
    });
  }

  return matches;
};

const App: React.FC<IProps> = ({ pageDatas, api }) => {
  const location = useLocation();
  const isFirstRenderRef = useRef(true);
  const matches = matchRoutes(routes, location.pathname);
  const [transformedMatches, setTransformedMatches] = useState(() =>
    transformMatches({ matches, pageDatas })
  );

  useEffect(() => {
    (async function () {
      if (!isFirstRenderRef.current) {
        const matches = matchRoutes(routes, location.pathname);

        const loadedPageDatas = await loadPageDatas({
          path: location.pathname,
          api,
        });

        setTransformedMatches(
          transformMatches({ matches, pageDatas: loadedPageDatas })
        );
      }

      isFirstRenderRef.current = false;
    })();
  }, [location]);

  return (
    <ApiContextProvider>{renderMatches(transformedMatches)}</ApiContextProvider>
  );
};

export default App;
