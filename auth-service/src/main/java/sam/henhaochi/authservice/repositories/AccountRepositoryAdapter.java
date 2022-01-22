package sam.henhaochi.authservice.repositories;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Repository;
import sam.henhaochi.authservice.constants.AccountCreationStatus;
import sam.henhaochi.authservice.constants.AccountLoginStatus;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.repositories.entities.AccountEntity;
import sam.henhaochi.authservice.repositories.mappers.AccountEntityMapper;
import sam.henhaochi.authservice.usecases.out.AccountDataSource;

import javax.annotation.Nullable;

@AllArgsConstructor
@Repository
public class AccountRepositoryAdapter implements AccountDataSource {

    final AccountRepository accountRepository;
    final AccountEntityMapper entityMapper;

    @Override
    public AccountCreationStatus registerWith(
            Account account
    ) {
        if (!accountRepository.existsByEmailEqualsIgnoreCase(
                account.getEmail()
        )){
            accountRepository.save(
                    entityMapper.map(
                            account
                    )
            );
            return AccountCreationStatus.SUCCESS;
        } else {
            return AccountCreationStatus.FAIL;
        }
    }

    @Nullable
    @Override
    public Account loginWith(Account account) {
        return accountRepository.login(
                account.getEmail(),
                account.getPassword(),
                account.getUsername()
        ).map(entityMapper::map).orElse(null);
    }
}
