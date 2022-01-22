package sam.henhaochi.authservice.usecases.in;

import sam.henhaochi.authservice.constants.AccountLoginStatus;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.repositories.entities.AccountEntity;

import java.security.NoSuchAlgorithmException;

public interface LoginAccountInput {
    Account with(
            Account account
    ) throws NoSuchAlgorithmException;
}
