UOW -> Unit of Work

{
    BEGIN
    Transacao 1 -> Repo 1
    Transacao 2 -> Repo 2
    COMMIT / ROLLBACK
}

-> Repositorios -> UOW -> getRepository -> Repositorio(TX)

Register
getRepositoryUnRegister
DO(fn func(uow uow) error) error {
    BEGIN
    Transacao 1 -> Repo 1
    Transacao 2 -> Repo 2
    COMMIT / ROLLBACK
}

fn(u)