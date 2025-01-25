create table produtos (
	id serial primary key,
	nome varchar(255),
	descricao varchar(255),
	preco decimal,
	quantidade integer
);

insert into produtos(nome, descricao, preco, quantidade)
values ('Camiseta', 'Preta', 19, 10), ('Fone', 'Muito bom', 99, 5);