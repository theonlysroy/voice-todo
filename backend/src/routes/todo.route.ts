import { Router } from "express";
import { todoController } from "../controllers/todo.controller";

const r = Router();

r.route("/").get(todoController.getAllTodos);
r.route("/").post(todoController.createTodo);
r.route("/:id/done").patch(todoController.markTodoDone);
r.route("/:id/remove").patch(todoController.deleteTodo);

export { r as todoRouter };
