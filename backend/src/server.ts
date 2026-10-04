import dotenv from "dotenv";
import express, {
  type Request,
  type NextFunction,
  type Response,
} from "express";
import connectDb from "./db";
import { todoRouter } from "./routes/todo.route";
dotenv.config({ path: ".env", quiet: true });

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (_, res) => {
  res.redirect("/healthz");
});

app.get("/healthz", (_, res: Response) => {
  res.status(200).json({ message: "api ok." });
});

// routes
app.use("/api/todos", todoRouter);

app.use((err: Error, req: Request, res: Response, next: NextFunction) => {
  if (res.headersSent) {
    return next(err);
  }
  console.log("error -->", err);
  return res.status(500).json({
    message: "Internal Server Error",
    errors: err.message,
  });
});

const srv = {
  addr: String(process.env.ADDR) || "http://127.0.0.1",
  port: parseInt(String(process.env.PORT) || "4000"),
  cb: () => console.log(`Server starting at ${srv.addr}:${srv.port}`),
};

async function startServer() {
  try {
    // connect db
    await connectDb();
    app.listen(srv.port, srv.addr, srv.cb);
  } catch (error) {
    console.log("Failed to start the server", error);
    process.exit(1);
  }
}

startServer();
