import mongoose from "mongoose";

export default async function connectDb() {
  try {
    const connectionString = String(process.env.MONGO_URL);
    if (!connectionString || connectionString === "") {
      console.log("db connection string required");
      throw new Error();
    }
    await mongoose.connect(connectionString, {
      serverSelectionTimeoutMS: 3000,
    });
    console.log("DB connected.");
  } catch (err) {
    throw new Error("Mongodb connection failed");
  }
}
