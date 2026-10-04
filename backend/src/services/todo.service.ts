import { Todo } from "../models/todo";

const findTodos = async () => {
  const allTodos = await Todo.find({ isDeleted: false }).lean();
  return allTodos;
};

const addTodo = async (payload: { title: string; description: string }) => {
  try {
    const newTodo = await Todo.create(payload);
    if (!newTodo) throw new Error("failed to create todo");
    return newTodo.toJSON();
  } catch (err) {
    throw err;
  }
};

const completeTodo = async (id: string) => {
  try {
    const todo = await Todo.findById(id);
    if (!todo) throw new Error("todo not found");
    todo.isCompleted = true;
    const updatedTodo = await todo.save();
    return updatedTodo.toJSON();
  } catch (err) {
    throw err;
  }
};

const removeTodo = async (id: string) => {
  try {
    const todo = await Todo.findById(id);
    if (!todo) throw new Error("todo not found");
    todo.isDeleted = true;
    const updatedTodo = await todo.save();
    if (!updatedTodo) throw new Error("failed to delete todo");
    return true;
  } catch (err) {
    throw err;
  }
};

export const todoService = {
  findTodos,
  addTodo,
  completeTodo,
  removeTodo,
};
