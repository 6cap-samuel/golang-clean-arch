package sam.henhaochi.authservice.usecases;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;
import sam.henhaochi.authservice.constants.AccountLoginStatus;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.usecases.in.LoginAccountInput;
import sam.henhaochi.authservice.usecases.out.AccountDataSource;

import java.security.NoSuchAlgorithmException;

@AllArgsConstructor
@Service
public class LoginAccountInteractor implements LoginAccountInput {

    final AccountDataSource accountDataSource;

    @Override
    public Account with(Account account) throws NoSuchAlgorithmException {
        account.encrypt();
        return accountDataSource.loginWith(account);
    }
}
