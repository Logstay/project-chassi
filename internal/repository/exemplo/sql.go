package exemplo

const qryInserirExemplo = `
INSERT INTO public.Exemplo
	(nome
		idade,
		telefone,
		endereco,
		Exemplo_cargo_id)
VALUES
	(
		:nome, 
		:idade
		:telefone
		:endereco
		:Exemplo_cargo_id) returning id
`
