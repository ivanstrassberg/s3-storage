module main

replace api => ./api

replace db => ./db

go 1.22.0

require api v0.0.0-00010101000000-000000000000

require db v0.0.0-00010101000000-000000000000
