/*
 * @Author: your name
 * @Date: 2020-12-28 23:33:40
 * @LastEditTime: 2021-01-20 19:41:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/middleware.go
 */
package handlers

// MiddlewareProductValidation will execute
import (
	"context"
	"net/http"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
)

// MiddlewareProductValidation called before the acutual handler called
// Just like Spring-AOP
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Error("Unable to deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Error("Unable to validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
