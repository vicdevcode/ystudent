import jwt from "jsonwebtoken";
import { access_token_secret } from "./config";

export interface JwtUser {
  id: string;
  email: string;
  role: string;
}

export const getPayload = (token: string) => {
  return jwt.verify(token, access_token_secret);
};

export const InstanceOfJwt = (object: any): object is JwtUser => {
  if ("id" in object && "email" in object && "role" in object) return true;
  return false;
};
