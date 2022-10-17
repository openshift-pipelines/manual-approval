/*
Copyright 2022 The OpenShift Pipelines Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-pipelines/manual-approval-gate/pkg/handlers"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	// FIXME use a real health check
	r.HandleFunc("/health", handlers.HealthCheck)
	// FIXME use a real readiness check
	r.HandleFunc("/readiness", handlers.HealthCheck)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
