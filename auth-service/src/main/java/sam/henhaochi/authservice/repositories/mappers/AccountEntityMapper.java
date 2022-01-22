package sam.henhaochi.authservice.repositories.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.repositories.entities.AccountEntity;

@Service
public class AccountEntityMapper {
    public AccountEntity map(Account account) {
         return AccountEntity.builder()
                 .email(account.getEmail())
                 .username(account.getUsername())
                 .password(account.getPassword())
                 .build();
    }

    public Account map(AccountEntity entity){
        return Account.builder()
                .email(entity.getEmail())
                .username(entity.getUsername())
                .password(entity.getPassword())
                .build();
    }
}
