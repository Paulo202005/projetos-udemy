insert into usuarios(nome, nick, email, senha)
values
("Usuario01", "usr01", "usuario01@gmail.com", "$2a$10$0FaB/w4ocX7SAZS7V0bzIuKtenLI0n6gbbh2twYGpO5rZdD5Qj0r."),
("Usuario02", "usr02", "usuario02@gmail.com", "$2a$10$0FaB/w4ocX7SAZS7V0bzIuKtenLI0n6gbbh2twYGpO5rZdD5Qj0r."),
("Usuario03", "usr03", "usuario03@gmail.com", "$2a$10$0FaB/w4ocX7SAZS7V0bzIuKtenLI0n6gbbh2twYGpO5rZdD5Qj0r.");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);