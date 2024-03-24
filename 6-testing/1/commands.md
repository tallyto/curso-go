# Executar o teste

```bash
go test
```

# Executar o teste no modo verboso

```bash
go test -v
```

# Cobertura de código

```bash
go test -coverprofile=coverage.out
```

# Cobertura de código em html

```bash
go test -coverprofile=coverage.out
```

# Executa um teste de benchmark
```bash
go test -bench=.
```

# Executa somente o benchmark ignorando os testes
```bash
go test -bench=. -run=^#
```

# Executa o benchmark 10x
```bash
go test -bench=. -run=^# -count=10
```

# Executa o benchmark por 3s
```bash
go test -bench=. -run=^# -count=10 -benchtime=3s
```