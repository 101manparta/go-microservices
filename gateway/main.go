package main

import (
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/users", proxy("http://localhost:8081/users"))
    http.HandleFunc("/products", proxy("http://localhost:8082/products"))

    log.Println("API Gateway running at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func proxy(target string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Proxy the request to target microservice
        req, err := http.NewRequest(r.Method, target, r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        req.Header = r.Header

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()

        // Copy response header
        for k, v := range resp.Header {
            for _, vv := range v {
                w.Header().Add(k, vv)
            }
        }
        w.WriteHeader(resp.StatusCode)
        io.Copy(w, resp.Body)
    }
}
