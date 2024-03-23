module github.com/tallyto/curso-go/5-packaging/3-go-mod-replace/sistema

go 1.22.1

replace github.com/tallyto/curso-go/5-packaging/3-go-mod-replace/math => ./math

require github.com/tallyto/curso-go/5-packaging/3-go-mod-replace/math v0.0.0-00010101000000-000000000000
