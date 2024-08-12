package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"template-api-pg/internal/models"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ExampleController(ctx context.Context, r *mux.Router, conn *sql.DB) {
	r.Path("").Methods(http.MethodGet).HandlerFunc(ExampleList(ctx, conn))
	r.Path("/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ExampleGet(ctx, conn))
	r.Path("").Methods(http.MethodPost).HandlerFunc(ExampleCreate(ctx, conn))
	r.Path("/{id:[0-9]+}").Methods(http.MethodPut).HandlerFunc(ExamplePut(ctx, conn))
	r.Path("/{id:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(ExampleDelete(ctx, conn))
	r.Path("").Methods(http.MethodOptions).HandlerFunc(ExampleOptions(ctx, conn))
}

func ExampleList(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		examples, err := models.Examples().All(ctx, conn)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to retrieve examples from database")
		}

		if examples == nil {
			examples = make(models.ExampleSlice, 0)
		}

		output, err := json.Marshal(examples)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to convert examples to json")
		}

		w.Write(output)
	}
}

func ExampleGet(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var id int

		if rawId, ok := vars["id"]; ok {
			convId, err := strconv.Atoi(rawId)
			if err != nil {
				ErrorResponse(w, err, 400, fmt.Sprintf("invalid id provided: %s", rawId))
				return
			}
			id = convId
		}

		exists, err := models.ExampleExists(ctx, conn, id)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to determine if example exists")
			return
		}
		if !exists {
			ErrorResponse(w, err, 404, fmt.Sprintf("example with id %d not found", id))
			return
		}

		example, err := models.FindExample(ctx, conn, id, boil.Infer().Cols...)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to retrieve example from database")
			return
		}

		output, err := json.Marshal(example)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to convert example to json")
		}

		w.Write(output)
	}
}

func ExampleCreate(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to read request body - %v", err))
			return
		}

		var example models.Example

		if err := json.Unmarshal(body, &example); err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to parse json from request body - %v", err))
			return
		}

		if err := example.Insert(r.Context(), conn, boil.Infer()); err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to insert new example - %v", err))
			return
		}

		res, err := json.Marshal(example)
		if err != nil {
			ErrorResponse(w, err, 500, fmt.Sprintf("failed to marshal created example for response - %v", err))
			return
		}

		w.Write(res)
		w.WriteHeader(201)
	}
}

func ExamplePut(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var id int

		if rawId, ok := vars["id"]; ok {
			convId, err := strconv.Atoi(rawId)
			if err != nil {
				ErrorResponse(w, err, 400, fmt.Sprintf("invalid id provided: %s", rawId))
				return
			}
			id = convId
		}

		exists, err := models.ExampleExists(ctx, conn, id)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to determine if example exists")
			return
		}
		if !exists {
			ErrorResponse(w, err, 404, fmt.Sprintf("example with id %d not found", id))
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to read request body - %v", err))
			return
		}

		columnsToUpdate, err := parsePutColumns(body)
		if err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to parse json from request body - %v", err))
			return
		}

		if len(columnsToUpdate.Cols) == 0 {
			ErrorResponse(w, err, 400, "no columns found to update example")
			return
		}

		var example models.Example

		if err := json.Unmarshal(body, &example); err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to parse json from request body - %v", err))
			return
		}

		example.ID = id

		if _, err := example.Update(ctx, conn, columnsToUpdate); err != nil {
			ErrorResponse(w, err, 400, fmt.Sprintf("failed to update example %d - %v", id, err))
			return
		}

		if err := example.Reload(ctx, conn); err != nil {
			ErrorResponse(w, err, 500, fmt.Sprintf("failed to refresh updated record - %v", err))
			return
		}

		res, err := json.Marshal(example)
		if err != nil {
			ErrorResponse(w, err, 500, fmt.Sprintf("failed to marshal created example for response - %v", err))
			return
		}

		w.Write(res)
		w.WriteHeader(201)

	}
}

func ExampleDelete(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var id int

		if rawId, ok := vars["id"]; ok {
			convId, err := strconv.Atoi(rawId)
			if err != nil {
				ErrorResponse(w, err, 400, fmt.Sprintf("invalid id provided: %s", rawId))
				return
			}
			id = convId
		}

		exists, err := models.ExampleExists(ctx, conn, id)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to determine if example exists")
			return
		}
		if !exists {
			ErrorResponse(w, err, 404, fmt.Sprintf("example with id %d not found", id))
			return
		}

		example, err := models.FindExample(ctx, conn, id, boil.Infer().Cols...)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to retrieve example from database")
			return
		}

		if _, err := example.Delete(ctx, conn); err != nil {
			ErrorResponse(w, err, 500, "failed to delete example from database")
			return
		}

		w.Write(nil)
	}
}

func ExampleOptions(ctx context.Context, conn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		example := models.Example{}

		output, err := json.Marshal(example)
		if err != nil {
			ErrorResponse(w, err, 500, "failed to convert examples to json")
		}

		w.Write(output)
	}
}
