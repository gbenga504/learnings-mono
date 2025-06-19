import React from "react";
import { Link } from "react-router-dom";

import { loadData } from "./loadData";

interface IProps {
  pageData: Awaited<ReturnType<typeof loadData>>;
}

export const Dashboard: React.FC<IProps> = ({ pageData }) => {
  return (
    <div>
      <h1>
        Hello, Dashboard Page
        <Link to="/">Go back</Link>
      </h1>
      {pageData.map((data) => (
        <li key={data.id}>{data.name}</li>
      ))}
    </div>
  );
};
