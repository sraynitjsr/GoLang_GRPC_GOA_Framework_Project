package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service Front End")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("sub", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/sub/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("divide", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/divide/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
