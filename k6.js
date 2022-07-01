import http from "k6/http";
import { check } from "k6";

export const options = {
  vus: 10,
  duration: "10s",
};

export default function () {
  const res = http.get("http://127.0.0.1:8080/proxy/aaaaa/hello");
  check(res, {
    "status code": (r) => r.status === 200,
    "verify resp text": (r) => r.body.includes("hello world"),
  });
}
