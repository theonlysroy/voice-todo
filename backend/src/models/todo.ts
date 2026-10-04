import { model, Schema } from "mongoose";

interface ITodo {
  title: string;
  description: string;
  isCompleted: boolean;
  isDeleted: boolean;
}

const todoSchema = new Schema<ITodo>(
  {
    title: { type: String, required: true, maxLength: 50 },
    description: { type: String, maxLength: 200 },
    isCompleted: { type: Boolean, default: false },
    isDeleted: { type: Boolean, default: false },
  },
  {
    timestamps: true,
  },
);

export const Todo = model<ITodo>("todo", todoSchema)
