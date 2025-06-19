import React, { useState } from "react";
import { useApi } from "../../api/ApiContext";
import { AxiosError } from "axios";
import { useNavigate } from "react-router-dom";

export const Login = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [inputsData, setInputsData] = useState({ email: "", password: "" });

  const handleChange = (changeset: Partial<typeof inputsData>) => {
    setInputsData((prevState) => ({
      ...prevState,
      ...changeset,
    }));
  };

  const handleSubmit = async (ev: React.MouseEvent) => {
    ev.preventDefault();

    try {
      const result = await api.auth.login({
        email: inputsData.email,
        password: inputsData.password,
      });

      if (result.data.status === "ok") {
        navigate("/dashboard");
      }
    } catch (error) {
      if (error instanceof AxiosError) {
        console.log("An error occurred man ===>", error);
      }
    }
  };

  return (
    <h1>
      Hello, Login Page
      <form
        method="post"
        style={{
          display: "flex",
          flexDirection: "column",
          rowGap: "5px",
          justifyContent: "center",
        }}
      >
        <input
          type="email"
          name="email"
          onChange={(ev) => handleChange({ email: ev.target.value })}
        />
        <input
          type="password"
          name="password"
          onChange={(ev) => handleChange({ password: ev.target.value })}
        />
        <button type="submit" onClick={handleSubmit}>
          Login
        </button>
      </form>
    </h1>
  );
};
