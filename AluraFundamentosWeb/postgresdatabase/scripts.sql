CREATE DATABASE aluraLojaDB;

CREATE TABLE produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

DELETE FROM produtos;

INSERT INTO produtos (nome, descricao, preco, quantidade)
values 
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 99, 55),
('Notebook', 'Muito rápido', 1999, 2),
('Tenis', 'Confortável', 89, 3),
('iPhone', '6s', 1800, 4);

SELECT id, nome, descricao, preco, quantidade FROM produtos