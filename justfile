generate flavor:
	@whiskers theme.hbs {{flavor}} > {{flavor}}.go
	@sed -i 's/hsl(\([0-9]*\), \([0-9]*\)%, \([0-9]*\)%)/\1, 0.\2, 0.\3/g' {{flavor}}.go
	@go fmt {{flavor}}.go

generate-all: (generate "latte") (generate "frappe") (generate "macchiato") (generate "mocha")
