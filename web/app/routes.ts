import { type RouteConfig, layout, route } from "@react-router/dev/routes";

export default [
  route("/", "./routes/login.tsx"),
  route("/legal", "./routes/legal.tsx"),

  layout("./routes/internal/layout.tsx", [
    route("/home", "./routes/internal/home.tsx"),
    route("/profile", "./routes/internal/profile.tsx"),
    route("/plans", "./routes/internal/plans.tsx"),
  ]),
] satisfies RouteConfig;
