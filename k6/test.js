import http from "k6/http";

export default () => {
    const res = http.get("http://localhost:8080/tests");
};
