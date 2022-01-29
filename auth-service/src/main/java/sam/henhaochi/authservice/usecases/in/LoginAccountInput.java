package sam.henhaochi.authservice.usecases.in;

import sam.henhaochi.authservice.entities.Account;

import java.security.NoSuchAlgorithmException;

public interface LoginAccountInput {
    Account with(
            Account account
    ) throws NoSuchAlgorithmException;
}
