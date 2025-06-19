import React from "react";

import { Login } from "./Login/Login";
import { Dashboard } from "./Dashboard/Dashboard";
import { loadData as dashboardLoadData } from "./Dashboard/loadData";

import { RouteObjectWithLoadData } from "react-router-dom";

export const routes: RouteObjectWithLoadData[] = [
  {
    id: "login",
    path: "/",
    component: Login,
  },
  {
    id: "dashboard",
    path: "/dashboard",
    component: Dashboard,
    loadData: dashboardLoadData,
  },
];
