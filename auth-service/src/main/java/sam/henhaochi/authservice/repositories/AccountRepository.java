package sam.henhaochi.authservice.repositories;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;
import sam.henhaochi.authservice.repositories.entities.AccountEntity;

import java.util.Optional;

@Repository
public interface AccountRepository
        extends CrudRepository<AccountEntity, String>{
    boolean existsByEmailEqualsIgnoreCase(String email);

    @Query("select a from Account a " +
            "where upper(a.email) = upper(?1) " +
            "and a.password = ?2 " +
            "or upper(a.username) = upper(?3) " +
            "and a.password = ?2")
    Optional<AccountEntity> login(String email, String password, String username);

}
