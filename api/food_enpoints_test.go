package api_test

import (
	"net/http"
	"net/http/httptest"
	"bytes"
	"fmt"


	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Foods API", func() {
	Describe("POST /api/v1/foods", func() {
		Context("with valid parameters", func() {
			It("should return the created food", func() {
				foodParams := []byte(`{"name":"Banana","calories":100}`)

				req, _ := http.NewRequest("POST", "/api/v1/foods", bytes.NewBuffer(foodParams))
				response := httptest.NewRecorder()

				app.Router.ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusCreated))
			})
		})
	})

	Describe("GET /api/v1/foods", func() {
		Context("with valid parameters", func() {
			It("should return the created food", func() {
				app.DB.Instance.QueryRow(
			    "INSERT INTO foods(name, calories) VALUES(Banana, 100)").Scan(&f.ID, &f.Name, &f.Calories)

				req, _ := http.NewRequest("GET", "/api/v1/foods", nil)
				response := httptest.NewRecorder()

				app.Router.ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusOK))
				Expect(repsonse.Body).To(Eqaul(`{"name":"Banana","calories":100}`))
			})
		})
	})
})
