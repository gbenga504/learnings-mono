import { Router } from "express";

const routes = Router();

routes.post("/auth/login", async function (req, res) {
  const { email, password } = req.body;
  const sevenDaysFromNow = new Date(Date.now() + 1000 * 60 * 60 * 24 * 7);

  const result = await req.api.auth.authenticate({ email, password });

  res.cookie("authToken", result.data.authToken, {
    httpOnly: true,
    secure: false,
    expires: sevenDaysFromNow,
  });

  res.status(200).json({ status: "ok" });
});

export default routes;
