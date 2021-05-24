module TCC/restapi

go 1.16

replace TCC/client => ./client

replace TCC/model => ./model

require (
	TCC/client v0.0.0-00010101000000-000000000000
	TCC/model v0.0.0-00010101000000-000000000000
	go.reizu.org/servemux v0.2.2
)
