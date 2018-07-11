package api_test

import (
	"net/http"
	"net/http/httptest"
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Foods API", func() {
	Describe("POST /api/v1/foods", func() {
		Context("with valid parameters", func() {
			It("should return the created food", func(done Done) {
				foodParams := []byte(`{"name":"Banana","calories":100}`)

				req, _ := http.NewRequest("POST", "/api/v1/foods", bytes.NewBuffer(foodParams))
				response := httptest.NewRecorder()
				app.Router.ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusCreated))
			})
		})
	})
})
