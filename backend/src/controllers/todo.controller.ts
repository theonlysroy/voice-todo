import asyncHandler from "../async";
import { todoService } from "../services/todo.service";

const getAllTodos = asyncHandler(async (req, res) => {
  const data = await todoService.findTodos();
  res.status(200).json({
    message: "all todos",
    data,
  });
});

const createTodo = asyncHandler(async (req, res) => {
  const { title, description } = req.body;
  if (!title)
    return res.status(400).json({
      message: "validation error",
      error: "title is required",
    });
  const data = await todoService.addTodo({ title, description });
  res.status(201).json({
    message: "todo created",
    data,
  });
});

const markTodoDone = asyncHandler(async (req, res) => {
  const { id } = req.params;
  if (!id)
    return res.status(400).json({
      message: "validation error",
      error: "todo id is required",
    });
  const data = await todoService.completeTodo(String(id));
  res.status(200).json({
    message: "todo marked completed.",
    data,
  });
});

const deleteTodo = asyncHandler(async (req, res) => {
  const { id } = req.params;
  if (!id)
    return res.status(400).json({
      message: "validation error",
      error: "todo id is required",
    });
  await todoService.removeTodo(String(id));
  return res.status(200).json({
    message: "todo deleted.",
  });
});

export const todoController = {
  getAllTodos,
  createTodo,
  markTodoDone,
  deleteTodo,
};
